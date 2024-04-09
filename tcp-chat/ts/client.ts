import { Socket } from "net"

export class Client {
   constructor(
      public connection: Socket,
      public name: string,
      public address: string
   ) { }

   quit() {
      this.connection.end()
   }
}
