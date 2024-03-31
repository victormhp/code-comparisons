export class Client {
   constructor(connection, name, address) {
      this.connection = connection
      this.name = name
      this.address = address
   }

   quit() {
      this.connection.end()
   }
}
