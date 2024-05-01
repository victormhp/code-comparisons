import fs from "node:fs"
import https from "node:https"
import readLine from "node:readline";
import { performance } from "node:perf_hooks";
import { type Readable } from "node:stream";

export function basename(s: string) {
    let slash = s.lastIndexOf("/")
    s = s.slice(slash + 1)
    let dot = s.lastIndexOf(".")
    if (dot >= 0) {
        s = s.slice(0, dot)
    }
    console.log(s)
}

export async function fetchUrl(urls: string[]): Promise<void> {
    urls.forEach(url => {
        https.get(url, res => {
            let data = "";
            res.on("data", chunk => data += chunk);
            res.on("end", () => {
                console.log(data);
                console.log("Status: ", res.statusCode);
            });
        }).on("error", err => {
            console.error(`fetch: ${err}`);
            process.exit(1);
        })
    })
}

export async function fetchAll(urls: string[]) {
    const start = performance.now();

    const fetchUrl = (url: string): Promise<string> => {
        const start = performance.now();
        return new Promise((resolve) => {
            https.get(url, res => {
                let data = "";
                res.on("data", chunk => data += chunk);
                res.on("end", () => {
                    const secs = (performance.now() - start) / 1000;
                    resolve(`${secs.toFixed(2)}s ${data.length} ${url}`);
                });
            }).on("error", err => {
                resolve(`fetch: ${url}: ${err.message}`);
            })
        })
    }

    const results = urls.map(url => fetchUrl(url));
    for (const r of results) {
        console.log(await r);
    }
    const secs = (performance.now() - start) / 1000;
    console.log(`${secs.toFixed(2)}s elapsed`);
}

function popCount(n: number): number {
    const pc = new Uint8Array(256);
    for (let i = 0; i < pc.length; i++) {
        pc[i] = pc[Math.floor(i / 2)] + (i & 1);
    }

    // JS bitwise operations are performed on 32-bit signed integers
    const low = n & 0xFFFFFFFF; // lower 32 bits
    const high = (n / 0x100000000) & 0xFFFFFFFF; // higher 32 bits

    return (pc[low & 0xFF] +
        pc[(low >> 8) & 0xFF] +
        pc[(low >> 16) & 0xFF] +
        pc[(low >> 24) & 0xFF] +
        pc[high & 0xFF] +
        pc[(high >> 8) & 0xFF] +
        pc[(high >> 16) & 0xFF] +
        pc[(high >> 24) & 0xFF]);
}

export function readPopCount(args: string[]): void {
    args.forEach(l => {
        l = l.trim();
        const num = Number(l);
        if (isNaN(num) || !Number.isInteger(num)) {
            console.error("Invalid Input");
            process.exit(1);
        }
        console.log(`${num} - Pop Count: ${popCount(num)}`);

    })
}
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
