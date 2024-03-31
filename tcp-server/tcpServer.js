const net = require("node:net")

const PORT = 8080
const server = net.createServer()

server.on('connection', socket => {
    socket.on('data', (data) => {
        console.log(`Data received from client: ${data}`);
        socket.write(`[*] Server echo: ${data}`);
    })

    socket.on('close', () => {
        console.log('[*] Client disconnected')
    })

    socket.on('error', (error) => {
        console.error(`[*] Connection error: ${error}`)
    })
})

server.listen(PORT, () => {
    console.log('[*] Server listening on', server.address())
})
