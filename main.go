package main

import "os"

func exiter(exitCode int) {
	os.Exit(exitCode)
}

func run(exit func(exitCode int)) {
	exit(0)
}

func main() {
	run(exiter)
}
