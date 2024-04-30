import concurrent.futures
import sys
import time
from typing import Dict, TextIO

import requests


def basename(s: str) -> str:
    slash = s.rfind("/")
    s = s[slash + 1:]
    if (dot := s.find(".")) >= 0:
        s = s[:dot]
    print(s)


def fetch(urls: list[str]):
    for url in urls:
        try:
            res = requests.get(url)
            print(res.text)
            print("Status: ", res.status)
        except requests.exceptions.HTTPError as e:
            print(f"fetch: HTTP error for {url}: {e}", file=sys.stderr)
            sys.exit(1)
        except requests.exceptions.RequestException as e:
            print(f"fetch: Error fetching {url}: {e}", file=sys.stderr)
            sys.exit(1)


def fetchall(urls: list[str]):

    def fetch_url(url: str):
        start = time.time()
        try:
            res = requests.get(url)
            nbytes = len(res.content)
            secs = time.time() - start
            return f"{secs:.2f}s {nbytes:7} {url}"
        except requests.exceptions.HTTPError as e:
            print(f"fetch: HTTP error for {url}: {e}", file=sys.stderr)
            sys.exit(1)
        except requests.exceptions.RequestException as e:
            print(f"fetch: Error fetching {url}: {e}", file=sys.stderr)
            sys.exit(1)

    start = time.time()
    with concurrent.futures.ThreadPoolExecutor() as executor:
        results = list(executor.map(fetch_url, urls))

    for r in results:
        print(r)

    print(f"{time.time() - start:.2f}s elapsed")


def popCount(n: int) -> int:
    pc = [0] * 256
    for i in range(256):
        pc[i] = pc[i // 2] + (i & 1)

    return (
        pc[n & 0xFF]
        + pc[(n >> 8) & 0xFF]
        + pc[(n >> 16) & 0xFF]
        + pc[(n >> 24) & 0xFF]
        + pc[(n >> 32) & 0xFF]
        + pc[(n >> 40) & 0xFF]
        + pc[(n >> 48) & 0xFF]
        + pc[(n >> 56) & 0xFF]
    )


def readPopCount(args: list[str]):

    for line in args:
        line = line.strip()
        try:
            num = int(line)
            print(f"{num} - Pop Count: {popCount(num)}")
        except ValueError as err:
            print(f"Invalid Input: {err}")


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
