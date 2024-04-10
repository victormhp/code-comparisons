import concurrent.futures
import sys
import time

import requests


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
