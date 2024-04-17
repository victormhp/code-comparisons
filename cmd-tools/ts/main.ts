import { uniq } from "./uniq";
import { readPopCount } from "./popcount"
import { fetchAll, fetchUrl } from "./fetch";

const [, , command, ...args] = process.argv;
if (!command) {
    console.error("Provide a command");
    process.exit(1);
}

const CMD = {
    uniq: "-uniq",
    fetch: "-fetch",
    fetchall: "-fetchall",
    popcount: "-popcount",
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
    default:
        console.error("Unknown command!");
        process.exit(1);
}
