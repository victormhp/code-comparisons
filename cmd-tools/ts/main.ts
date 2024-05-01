import { uniq, fetchUrl, fetchAll, readPopCount, basename } from "./utils";

const [command, ...args] = process.argv.slice(2);
if (!command) {
    console.error("Provide a command");
    process.exit(1);
}

const CMD = {
    uniq: "-uniq",
    fetch: "-fetch",
    fetchall: "-fetchall",
    popcount: "-popcount",
    basename: "-basename",
}

switch (command) {
    case CMD.uniq:
        uniq(args);
        break;
    case CMD.fetch:
        fetchUrl(args);
        break;
    case CMD.fetchall:
        fetchAll(args);
        break;
    case CMD.popcount:
        readPopCount(args);
        break;
    case CMD.basename:
        basename(args[0]);
        break;
    default:
        console.error("Unknown command!");
        process.exit(1);
}
