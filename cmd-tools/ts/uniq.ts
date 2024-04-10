import fs from "node:fs"
import readLine from "node:readline";
import { type Readable } from "node:stream";

export async function uniq(files: string[]): Promise<void> {
    let counts: Map<string, number> = new Map();
    if (files.length == 0) {
        await countLines(process.stdin, counts);
    } else {
        for (let f of files) {
            try {
                const stream = fs.createReadStream(f);
                await countLines(stream, counts);
                stream.close();
            } catch (err) {
                console.error("uniq-js", err);
            }
        }
    }

    for (let [line, n] of counts) {
        if (n > 1) {
            console.log(`${n}\t${line}`);
        }
    }
}

async function countLines(stream: Readable, counts: Map<string, number>) {
    const rl = readLine.createInterface({
        input: stream,
    })

    for await (const line of rl) {
        counts.set(line, (counts.get(line) || 0) + 1);
    }
}
