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
