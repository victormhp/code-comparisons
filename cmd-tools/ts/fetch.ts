import https from "node:https"
import { performance } from "node:perf_hooks";

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
