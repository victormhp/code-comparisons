import socket


class Client:

    def __init__(self, connection: socket.socket, name: str, address: str):
        self.connection = connection
        self.name = name
        self.address = address

    def quit(self):
        self.connection.close()
