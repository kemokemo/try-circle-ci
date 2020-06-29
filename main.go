package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

var (
	out     io.Writer = os.Stdout
	outerr  io.Writer = os.Stderr
	exit              = os.Exit
	cmdArgs []string  // commad line args hack

	ver bool
)

func init() {
	testing.Init() // require Go 1.13 or later
	flag.BoolVar(&ver, "v", false, "display version")
	flag.BoolVar(&ver, "version", false, "display version")
	flag.Parse()
	cmdArgs = flag.Args()
}

func main() {
	exit(run(cmdArgs))
}

func run(args []string) int {
	if ver {
		fmt.Fprintf(out, "%s version %s.%s", Name, Version, revision)
		return exitCodeOK
	}

	if len(args) == 0 {
		fmt.Fprintln(outerr, "please specify args")
		return exitCodeFailed
	}

	for _, arg := range args {
		g, err := greeter(arg)
		if err != nil {
			fmt.Fprintln(outerr, "failed to execute the greeter:", err)
			return exitCodeFailed
		}
		fmt.Fprintln(out, g)
	}

	return exitCodeOK
}
