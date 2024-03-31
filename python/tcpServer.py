import socket
import threading

HOST = "127.0.0.1"
PORT = 8080


def main():
    try:
        server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server.bind((HOST, PORT))
        server.listen()
        print(f"[*] Server listening on {HOST}:{PORT}")
    except socket.error as err:
        print(f"Failed to start server: {err}")
        return

    while True:
        try:
            connection, address = server.accept()
        except socket.error as err:
            print(f"Error accepting connection: {err}")
            continue

        print(f"[*] Connection accepted from {address}")
        connection_handler = threading.Thread(
            target=handle_connection, args=(connection,)
        )
        connection_handler.start()


def handle_connection(connection: socket.socket):
    text = connection.recv(1024)
    message = text.decode("utf8").strip()
    print(f"[*] Received: {message}")
    connection.close()


if __name__ == "__main__":
    main()
