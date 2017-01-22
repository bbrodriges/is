# is

[![Build Status](https://travis-ci.org/bbrodriges/is.svg?branch=master)](https://travis-ci.org/bbrodriges/is)
[![GoDoc](https://godoc.org/github.com/bbrodriges/is?status.svg)](https://godoc.org/github.com/bbrodriges/is)
[![Go Report Card](https://goreportcard.com/badge/github.com/bbrodriges/is)](https://goreportcard.com/report/github.com/bbrodriges/is)

Micro check library in Golang. 

## installation

`go get github.com/bbrodriges/is`

## No regexs as much as possible

It works with runes as long as is it possible.
Part of source code:

```go
// Alpha check if the string contains only letters (a-zA-Z).
func Alpha(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') {
			return false
		}
	}
	return true
}
```

## usage

```go
package main

import "github.com/bbrodriges/is"

func main()  {
    is.Email("jhon@example.com") // true
    is.Numeric("Ⅸ") // false
    is.UTFNumeric("Ⅸ") // true
}
```

for more documentation [godoc](https://godoc.org/github.com/bbrodriges/is)

## Contribute

- Report problems
- Add/Suggest new features/recipes
- Improve/fix documentation

## Thanks & Authors

Original idea by

- [alioygur/is](https://github.com/alioygur/is) Micro check library