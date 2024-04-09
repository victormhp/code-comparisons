import sys

from fetch import fetch, fetchall
from uniq import uniq

args = sys.argv
command = args[1]
params = args[2:]

match command:
    case "-uniq":
        uniq(params)
    case "-fetch":
        fetch(params)
    case "-fetchall":
        fetchall(params)
    case _:
        print("Please provide a command")
        sys.exit(1)
