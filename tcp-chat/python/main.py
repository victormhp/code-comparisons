import logging
import socket
import threading

from client import Client
from message import Message

logging.basicConfig(level=logging.INFO, format="%(asctime)s - %(message)s")

HOST = "127.0.0.1"
PORT = 8080
CMD_QUIT = "/quit"

clients: list[Client] = []


def main():
    try:
        server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server.bind((HOST, PORT))
        server.listen()
        logging.info(f"Server listening on {HOST}:{PORT}")
    except socket.error as err:
        logging.error(f"Failed to start server: {err}")
        return

    while True:
        try:
            connection, address = server.accept()
        except socket.error as e:
            logging.error(f"Error accepting connection: {e}")
            break

        logging.info(f"{address} has connected")

        # Get client's name and register it
        connection.send(b"Enter your name: ")
        name = connection.recv(1024).decode("utf-8").strip()
        newClient = Client(connection, name, address)
        clients.append(newClient)

        # Welcome to the server
        connection.send(f"\nWelcome to the server {name}!\n".encode("utf-8"))
        connection.send(
            f"Type any message to send it, type {CMD_QUIT} to finish\n\n".encode(
                "utf-8"
            )
        )

        connection_handler = threading.Thread(
            target=handleConnection, args=(newClient,)
        )
        connection_handler.start()


def handleConnection(client: Client):
    broadcast(f"{client.name} joined the chat\n", client)
    while True:
        try:
            text = client.connection.recv(1024).decode("utf-8").strip()
            message = Message(client, text)
            parseMessage(message)
        except OSError:
            if client in clients:
                clients.remove(client)
                client.quit()
            break


def parseMessage(message: Message):
    if message.message == CMD_QUIT:
        broadcast(f"{message.client.name} leaved the chat\n", message.client)
        logging.info(f"{message.client.address} has disconnected")
        message.client.quit()
        clients.remove(message.client)
    else:
        message.client.connection.send(clearLine().encode("utf-8"))
        broadcast(message.string())


def broadcast(message: str, excludedClient: Client = None):
    for client in clients:
        if client != excludedClient:
            client.connection.send(message.encode("utf8"))


def clearLine():
    return "\x1B[A\x1B[2K"


if __name__ == "__main__":
    main()
