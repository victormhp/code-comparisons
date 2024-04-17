function popCount(n: number): number {
    const pc = new Uint8Array(256);
    for (let i = 0; i < pc.length; i++) {
        pc[i] = pc[Math.floor(i / 2)] + (i & 1);
    }

    // JS bitwise operations are performed on 32-bit signed integers
    const low = n & 0xFFFFFFFF; // lower 32 bits
    const high = (n / 0x100000000) & 0xFFFFFFFF; // higher 32 bits

    return (pc[low & 0xFF] +
        pc[(low >> 8) & 0xFF] +
        pc[(low >> 16) & 0xFF] +
        pc[(low >> 24) & 0xFF] +
        pc[high & 0xFF] +
        pc[(high >> 8) & 0xFF] +
        pc[(high >> 16) & 0xFF] +
        pc[(high >> 24) & 0xFF]);
}

export function readPopCount(args: string[]): void {
    args.forEach(l => {
        l = l.trim();
        const num = Number(l);
        if (isNaN(num) || !Number.isInteger(num)) {
            console.error("Invalid Input");
            process.exit(1);
        }
        console.log(`${num} - Pop Count: ${popCount(num)}`);

    })
}
