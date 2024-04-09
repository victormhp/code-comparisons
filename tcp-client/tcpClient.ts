import net from "node:net"
import readline from "node:readline"

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
})

const HOST = 'localhost'
const PORT = 8080

const client = new net.Socket()

client.connect(PORT, HOST, () => {
    console.log('Connected to the server')
    rl.question("Message: ", (msg: string) => {
        client.write(msg)
        rl.close()
    })
})

client.on('data', (data: Buffer) => {
    console.log(`Received from server: ${data}`)
    client.destroy()
})

client.on('close', () => {
    console.log('Connection closed')
})

client.on('error', (error: Error) => {
    console.error("Error:", error)
})
