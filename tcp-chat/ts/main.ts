import net from "node:net"
import { Client } from "./client.ts"
import { Message } from "./message.ts"

const CMD_QUIT = "/quit"
const PORT = 8080

let clients: Client[] = []
const server = net.createServer()

server.on('connection', (socket: net.Socket) => {
   let client: Client | null = null;
   let address = '';
   if (socket.remoteAddress) {
      address = address.replace(/^::ffff:/, '') + ":" + socket.remotePort
   }
   console.log(`[*] ${address} has connected`)
   socket.write("Enter your username: ")

   socket.on('data', (data) => {
      const text = data.toString().trim()
      if (client == null) {
         client = new Client(socket, text, address)
         clients.push(client)
         broadcast(`${client.name} joined the chat!\n`, socket)
         socket.write(`\nWelcome ${client.name}!\n`)
         socket.write(`Type any message to send it, type ${CMD_QUIT} to finish\n\n`)
      } else {
         const message = new Message(client, text)
         parseMessage(message)
      }
   })

   socket.on('close', () => {
      console.log(`[*] ${address} has disconnected`)
   })

   socket.on('error', (err: Error) => {
      console.error("[*] Error:", err)
      clients = clients.filter(c => c !== client)
      socket.end()
   })
})

server.listen(PORT, () => {
   console.log(`[*] Server listening on :${PORT}`)
})

function parseMessage(message: Message) {
   const name = message.client.name
   const socket = message.client.connection

   if (message.message === CMD_QUIT) {
      broadcast(`${name} leaved the chat!\n`, socket)
      clients = clients.filter(c => c !== message.client)
      socket.end()
   } else {
      clearLine(socket)
      broadcast(message.string())
   }
}

function broadcast(message: string, excludeSocket: net.Socket | null = null) {
   for (let client of clients) {
      if (client.name !== null && client.connection !== excludeSocket) {
         client.connection.write(message)
      }
   }
}

function clearLine(socket: net.Socket) {
   // \x1B[A moves the cursor up one line.
   // \x1B[2K clears the entire line where the cursor is currently located.
   socket.write("\x1B[A\x1B[2K")
}
