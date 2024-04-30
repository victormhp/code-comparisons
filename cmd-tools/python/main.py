import argparse
import sys

from utils import basename, fetch, fetchall, readPopCount, uniq

parser = argparse.ArgumentParser(description="Command Line Tool")
parser.add_argument("-command", required=True, help="Command to run")
parser.add_argument("args", nargs="*", help="Arguments for the command")
args = parser.parse_args()
if not args:
    print("Provide a command")
    sys.exit(1)

match args.command:
    case "uniq":
        uniq(args.args)
    case "fetch":
        fetch(args.args)
    case "fetchall":
        fetchall(args.args)
    case "popcount":
        readPopCount(args.args)
    case "basename":
        basename(args.args[0])
    case _:
        print(f"Unknown tool: {args.command}")
        sys.exit(1)
