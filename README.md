# Cross-platform access to system clipboard

[![GoDoc](https://godoc.org/github.com/ardnew/pasta?status.svg)](http://godoc.org/github.com/ardnew/pasta)

Go module and CLI utility that works on Windows, WSL, macOS and Linux.

## Install

```bash
# Install the latest CLI utility only
go install github.com/ardnew/pasta/cmd/pasta@latest

# Or, add latest version as dependency to your Go module
go get github.com/ardnew/pasta
```

## Usage

- CLI
    ```bash
    $ pasta -h
    USAGE
      pasta -           # (copy) copy stdin to clipboard
      pasta             # (paste) write clipboard contents to stdout

      pasta < FILE      # (copy) write FILE contents to clipboard
      pasta > FILE      # (paste) overwrite FILE with clipboard contents

      COMMAND | pasta   # (copy) write COMMAND output to clipboard
      pasta | COMMAND   # (paste) pipe clipboard contents to COMMAND
    ```

- Module
  - See [godoc](http://godoc.org/github.com/ardnew/pasta) for detailed API.
  - Copy to clipboard
    ```go
    clipboard.WriteAll("Hello World!")
    ```
  - Paste from clipboard
    ```go
    text, err := clipboard.ReadAll()
    ```

## License

[MIT](LICENSE)
