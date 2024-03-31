const net = require("node:net")
const readline = require("readline")

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
})

const HOST = 'localhost'
const PORT = 8080

const client = new net.Socket()

client.connect(PORT, HOST, () => {
    console.log('Connected to the server')
    rl.question("Message: ", (msg) => {
        client.write(msg)
        rl.close()
    })
})

client.on('data', (data) => {
    console.log(`Received from server: ${data}`)
    client.destroy()
})

client.on('close', () => {
    console.log('Connection closed')
})

client.on('error', (error) => {
    console.error("Error:", error)
})
