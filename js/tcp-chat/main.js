import net from "node:net"

const CMD_QUIT = "/quit"

const clients = new Map()
const PORT = 8080
const server = net.createServer()

function clearLine(socket) {
   // \x1B[A moves the cursor up one line.
   // \x1B[2K clears the entire line where the cursor is currently located.
   socket.write("\x1B[A\x1B[2K")
}

function broadcast(message, excludeSocket = null) {
   for (let [socket, client] of clients) {
      if (client.name !== null && socket !== excludeSocket) {
         socket.write(message)
      }
   }
}

function parseMessage(socket, data) {
   const client = clients.get(socket)

   if (client.name === null) {
      client.name = data
      socket.write(`\nWelcome ${client.name}!\n`)
      socket.write(`Type any message to send it, type ${CMD_QUIT} to finish\n\n`)
      broadcast(`${client.name} joined the chat!\n`, socket)
   } else if (data === CMD_QUIT) {
      broadcast(`${client.name} leaved the chat!\n`, socket)
      clients.delete(socket)
      socket.end()
   } else {
      const message = `[${client.name}] -> ${data}\n`
      clearLine(socket)
      broadcast(message)
   }
}

server.on('connection', (socket) => {
   socket.write("Enter your username: ")

   const address = socket.remoteAddress.replace(/^::ffff:/, '') + ":" + socket.remotePort
   console.log(`[*] ${address} has connected`)
   clients.set(socket, { name: null, address })

   socket.on('data', (data) => {
      data = data.toString().trim()
      parseMessage(socket, data)
   })

   socket.on('close', () => {
      console.log(`[*] ${address} has disconnected`)
   })

   socket.on('error', (err) => {
      console.error("[*] Error:", err)
      clients.delete(socket)
      socket.end()
   })
})

server.listen(PORT, () => {
   console.log(`[*] Server listening on :${PORT}`)
})
