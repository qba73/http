[![Go](https://github.com/qba73/http/actions/workflows/go.yml/badge.svg)](https://github.com/qba73/http/actions/workflows/go.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/http?logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/http)](https://goreportcard.com/report/github.com/qba73/http)

# http

```http``` is a Go library that provides a drop-in replacement functionality for running the http server configured to use TLS.

## Using the Go library

Import the library using:
```go
import httpb "github.com/qba73/http"
```

Start a default http server with TLS:
```go
err := httpb.ListenAndServeTLS(":443", cert, key, handler)
if err != nil {
    log.Fatalln(err)
}
```

The function signature is the key difference compared with the same func from the http standard library. The function takes a cert and a key which are slices of bytes (```[]byte```), not paths to files located in the filesystem!

Using ```[]byte``` types instead of paths allows passing a cert and a key directly to the function without creating a temporary directory to store the content of the cert and key in files in the filesystem.
