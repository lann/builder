# Builder - fluent immutable builders

[![Build Status](https://travis-ci.org/lann/builder.png?branch=master)](https://travis-ci.org/lann/builder)

[![GoDoc](https://godoc.org/github.com/lann/builder?status.png)](https://godoc.org/github.com/lann/builder)

## fluent
```go
req := ReqBuilder.
Url("http://golang.org").
Header("User-Agent", "Builder").
Get()
```

## immutable
```go
build := WordBuilder.Letters("Build")
builder := build.Letters("er").String()
building := build.Letters("ing").String()
```

## builders
```go
muppet := MuppetBuilder.
Name("Beaker").
Hair("orange").
GetStruct()
muppet.Name == "Beaker"
```
