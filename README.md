# is

[![Build Status](https://travis-ci.org/bbrodriges/is.svg?branch=master)](https://travis-ci.org/bbrodriges/is)
[![GoDoc](https://godoc.org/github.com/bbrodriges/is?status.svg)](https://godoc.org/github.com/bbrodriges/is)
[![Go Report Card](https://goreportcard.com/badge/github.com/bbrodriges/is)](https://goreportcard.com/report/github.com/bbrodriges/is)
[![Coverage Status](https://coveralls.io/repos/github/bbrodriges/is/badge.svg?branch=master)](https://coveralls.io/github/bbrodriges/is?branch=master)

Micro check library in Golang. 

## Installation

`go get github.com/bbrodriges/is`

## No regexs as much as possible

It works with runes, basic operations (loops, conditions etc) and standart library functions as long as is it possible.
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

## Usage

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

## Benchmarks

If application speed is important to you as for me, you can check out benchmarks in [every Travis CI build](https://travis-ci.org/bbrodriges/is).

## Contribute

- Report problems
- Add/Suggest new features/recipes
- Improve/fix documentation

## Thanks & Authors

Original project:

- [alioygur/is](https://github.com/alioygur/is) Micro check library

## Main differences

This fork has been created because of:

- Some fundamental [disagreements](https://github.com/alioygur/is/issues/6) with author of original repo
- Unpleasant amount of regular expressions in code.

One of main goals of this repo is to fix all of the above claims.