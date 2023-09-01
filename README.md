# Clipboard for Go

[![GoDoc](https://godoc.org/github.com/antonmedv/clipboard?status.svg)](http://godoc.org/github.com/antonmedv/clipboard)

Provide copying and pasting to the Clipboard for Go.

```
go get github.com/antonmedv/clipboard
```

Supported platforms:
- macOS
- Windows + WSL
- Linux

## Usage

```go
import "github.com/antonmedv/clipboard"

clipboard.WriteAll("Hello World!")

text, err := clipboard.ReadAll()
```

## License

[MIT](LICENSE)
