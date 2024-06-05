package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ardnew/pasta"
	"github.com/pkg/errors"
)

const version = "1.2.1"

var exe string

func init() {
	var err error
	exe, err = os.Executable()
	if err != nil {
		exe = "%!s(BADEXE)"
	} else {
		exe = filepath.Base(exe)
	}
}

func usage() {
	desc := `%[1]s version %[2]s usage:

  %[1]s -           # (copy) copy stdin to clipboard
  %[1]s             # (paste) write clipboard contents to stdout

  %[1]s < FILE      # (copy) write FILE contents to clipboard
  %[1]s > FILE      # (paste) overwrite FILE with clipboard contents

  COMMAND | %[1]s   # (copy) write COMMAND output to clipboard
  %[1]s | COMMAND   # (paste) pipe clipboard contents to COMMAND

`
	fmt.Printf(desc, exe, version)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	assertNil := func(err error) {
		if err != nil {
			log.Fatalf("error: %v\n\t(see: %s -h)", err, exe)
		}
	}

	if !isInteractive() || flag.Arg(0) == "-" {
		// copy mode: stdin is connected to a pipe or file descriptor
		out, err := io.ReadAll(os.Stdin)
		assertNil(errors.Wrap(err, "read from stdin"))
		err = pasta.WriteAll(string(out))
		assertNil(errors.Wrap(err, "write to clipboard"))
	} else {
		// paste mode: stdin is not in a valid state for reading
		out, err := pasta.ReadAll()
		assertNil(errors.Wrap(err, "read from clipboard"))
		_, err = os.Stdout.WriteString(out)
		assertNil(errors.Wrap(err, "write to stdout"))
	}
}
