import sys
from typing import Dict, TextIO


def uniq(files: list[str]):
    counts = {}
    if not files:
        countLines(sys.stdin, counts)
    else:
        for arg in files:
            try:
                with open(arg, "r") as f:
                    countLines(f, counts)
            except IOError as err:
                print(f"uniq-py: {err}", file=sys.stderr)

        for line, n in counts.items():
            if n > 1:
                print(f"{n}\t{line}")


def countLines(f: TextIO, counts: Dict[str, int]):
    for line in f:
        line = line.strip()
        counts[line] = counts.get(line, 0) + 1
