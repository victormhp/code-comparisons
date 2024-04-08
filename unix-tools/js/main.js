import { uniq } from "./uniq.js";

const [, , command, ...flags] = process.argv;
if (!command) {
    console.error("Provide a unix command");
    process.exit(1);
}

const CMD_QUIT = "-uniq";

switch (command) {
    case CMD_QUIT:
        uniq(flags).catch(err => console.error(err));
        break;
    default:
        console.error("Unknown command!");
        process.exit(1);
}
