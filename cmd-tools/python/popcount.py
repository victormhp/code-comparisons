def popCount(n: int) -> int:
    pc = [0] * 256
    for i in range(256):
        pc[i] = pc[i // 2] + (i & 1)

    return (pc[n & 0xFF] + pc[(n >> 8) & 0xFF] + pc[(n >> 16) & 0xFF] +
            pc[(n >> 24) & 0xFF] + pc[(n >> 32) & 0xFF] +
            pc[(n >> 40) & 0xFF] + pc[(n >> 48) & 0xFF] + pc[(n >> 56) & 0xFF])


def readPopCount(args: list[str]):

    for line in args:
        line = line.strip()
        try:
            num = int(line)
            print(f"{num} - Pop Count: {popCount(num)}")
        except ValueError as err:
            print(f"Invalid Input: {err}")
