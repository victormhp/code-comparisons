// Get command line arguments
function printArgs() {
    const args = process.argv

    args.forEach((a, i) => {
        console.log(i, "->", a)
    })
}
