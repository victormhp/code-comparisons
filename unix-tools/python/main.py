import sys

from uniq import uniq

args = sys.argv
command = args[1]
files = args[2:]

match command:
    case "-uniq":
        uniq(files)
    case _:
        print("Please provide a command")
        sys.exit(1)
