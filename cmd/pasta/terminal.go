//go:build !windows
// +build !windows

package main

import (
	"os"

	"github.com/mattn/go-isatty"
)

func isInteractive() bool {
	return isatty.IsTerminal(os.Stdin.Fd())
}
