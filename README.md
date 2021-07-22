Hyperscale secure [![Last release](https://img.shields.io/github/release/hyperscale-stack/secure.svg)](https://github.com/hyperscale-stack/secure/releases/latest) [![Documentation](https://godoc.org/github.com/hyperscale-stack/secure?status.svg)](https://godoc.org/github.com/hyperscale-stack/secure)
====================

[![Go Report Card](https://goreportcard.com/badge/github.com/hyperscale-stack/secure)](https://goreportcard.com/report/github.com/hyperscale-stack/secure)

| Branch  | Status | Coverage |
|---------|--------|----------|
| master  | [![Build Status](https://github.com/hyperscale-stack/secure/workflows/Go/badge.svg?branch=master)](https://github.com/hyperscale-stack/secure/actions?query=workflow%3AGo) | [![Coveralls](https://img.shields.io/coveralls/hyperscale-stack/secure/master.svg)](https://coveralls.io/github/hyperscale-stack/secure?branch=master) |

The Hyperscale secure is a secure string generator.

## Example

```go
package main

import (
    "fmt"
    "github.com/hyperscale-stack/secure"
)

func main() {
    s, err := secure.GenerateRandomString(64)
    if err != nil {
        panic(err)
    }

    fmt.Println(s)
}

```

## License

Hyperscale secure is licensed under [the MIT license](LICENSE.md).
