go getlogin
===========

Package getlogin provides functionality similar libc's getlogin(3).

This package is provided as a convience and should NOT be used for any security related functionality. Just like libc's getlogin(3), this package can be easily tricked into returning incorrect information.

## Documentation

https://godoc.org/github.com/psanford/getlogin

## Example

    package main

    import (
    	"fmt"

    	"github.com/psanford/getlogin"
    )

    func main() {
    	fmt.Println("getlogin: ", getlogin.GetLogin())
    }

## Installation

    go get github.com/psanford/getlogin
