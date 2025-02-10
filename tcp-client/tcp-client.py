import socket

HOST = "127.0.0.1"
PORT = 8080

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect((HOST, PORT))

try:
    message = input("Message: ")
    client.sendall(message.encode("utf8"))
finally:
    print("Exiting server")
    client.close()
