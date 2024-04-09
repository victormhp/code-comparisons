import { Client } from "./client"

export class Message {
   constructor(
      public client: Client,
      public message: string
   ) { }

   string() {
      return `[${this.client.name}] -> ${this.message}\n`
   }
}
