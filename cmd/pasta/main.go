package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ardnew/pasta"
)

func usage() {
	exe, err := os.Executable()
	if err != nil {
		exe = "%!s(BADEXECUTABLE)"
	} else {
		exe = filepath.Base(exe)
	}
	desc := `USAGE
	%[1]s -           # (copy) copy stdin to clipboard
	%[1]s             # (paste) write clipboard contents to stdout

	%[1]s < FILE      # (copy) write FILE contents to clipboard
	%[1]s > FILE      # (paste) overwrite FILE with clipboard contents

	COMMAND | %[1]s   # (copy) write COMMAND output to clipboard
	%[1]s | COMMAND   # (paste) pipe clipboard contents to COMMAND
`
	fmt.Printf(desc, exe)
}

func assert[T comparable](v, x T) {
	if v != x {
		panic(fmt.Errorf("assertion failed: %v != %v", v, x))
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if !isInteractive() || flag.Arg(0) == "-" {
		// copy mode: stdin is connected to a pipe or file descriptor
		out, err := io.ReadAll(os.Stdin)
		assert(err, nil)
		err = pasta.WriteAll(string(out))
		assert(err, nil)
	} else {
		// paste mode: stdin is not in a valid state for reading
		out, err := pasta.ReadAll()
		assert(err, nil)
		_, err = os.Stdout.WriteString(out)
		assert(err, nil)
	}
}
