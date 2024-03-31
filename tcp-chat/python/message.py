from client import Client


class Message:

    def __init__(self, client: Client, message: str):
        self.client = client
        self.message = message

    def string(self):
        return f"[{self.client.name}] -> {self.message}\n"
