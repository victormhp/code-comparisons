import net from "node:net"

const PORT = 8080
const server = net.createServer()

server.on('connection', (socket: net.Socket) => {
    socket.on('data', (data: Buffer) => {
        console.log(`Data received from client: ${data}`);
        socket.write(`[*] Server echo: ${data}`);
    })

    socket.on('close', () => {
        console.log('[*] Client disconnected')
    })

    socket.on('error', (error: Error) => {
        console.error(`[*] Connection error: ${error}`)
    })
})

server.listen(PORT, () => {
    console.log('[*] Server listening on', server.address())
})
