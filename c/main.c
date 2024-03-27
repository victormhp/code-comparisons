#include <stdio.h>

void printArgs(int argc, char* argv[])
{
    for (int i = 0; i < argc; i++) {
        printf("%d -> %s\n", i, argv[i]);
    }
}

int main(int argc, char* argv[])
{
}
