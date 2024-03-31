export class Message {
   constructor(client, message) {
      this.client = client
      this.message = message
   }

   string() {
      return `[${this.client.name}] -> ${this.message}\n`
   }
}
