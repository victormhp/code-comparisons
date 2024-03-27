import sys


# Get command line arguments
def printArgs():
    args = sys.argv
    for i, v in enumerate(args):
        print(f"{i} -> {v}")
