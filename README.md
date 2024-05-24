# Cross-platform access to system clipboard

[![GoDoc](https://godoc.org/github.com/ardnew/pasta?status.svg)](http://godoc.org/github.com/ardnew/pasta)

[Go module](.) and [CLI utility](cmd/pasta) that works on Windows, WSL, macOS and Linux.

See [cmd/pasta](cmd/pasta) for installing and using the CLI utility.

## Install

```bash
go get github.com/ardnew/pasta
```

## Usage

See [godoc](http://godoc.org/github.com/ardnew/pasta) for detailed API.

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
