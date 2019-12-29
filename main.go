package main

import (
	"fmt"
	"os"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	for _, arg := range args {
		g, err := greeter(arg)
		if err != nil {
			fmt.Println("failed to execute the greeter:", err)
			return exitCodeFailed
		}
		fmt.Println(g)
	}

	return exitCodeOK
}
