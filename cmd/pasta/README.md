# Cross-platform access to system clipboard

[![GoDoc](https://godoc.org/github.com/ardnew/pasta?status.svg)](http://godoc.org/github.com/ardnew/pasta)

Go module and CLI utility that works on Windows, WSL, macOS and Linux.

## Install

```bash
go install github.com/ardnew/pasta/cmd/pasta@latest
```

## Usage

```bash
$ pasta -h
usage:
  pasta -           # (copy) copy stdin to clipboard
  pasta             # (paste) write clipboard contents to stdout
  pasta < FILE      # (copy) write FILE contents to clipboard
  pasta > FILE      # (paste) overwrite FILE with clipboard contents
  COMMAND | pasta   # (copy) write COMMAND output to clipboard
  pasta | COMMAND   # (paste) pipe clipboard contents to COMMAND
```

## License

[MIT](LICENSE)
