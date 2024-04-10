import { uniq } from "./uniq";
import { fetchAll, fetchUrl } from "./fetch";

const [, , command, ...flags] = process.argv;
if (!command) {
    console.error("Provide a unix command");
    process.exit(1);
}

const CMD = {
    uniq: "-uniq",
    fetch: "-fetch",
    fetchall: "-fetchall",
}

switch (command) {
    case CMD.uniq:
        uniq(flags);
        break;
    case CMD.fetch:
        fetchUrl(flags);
        break;
    case CMD.fetchall:
        fetchAll(flags);
        break;
    default:
        console.error("Unknown command!");
        process.exit(1);
}
