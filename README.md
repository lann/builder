# Builder - fluent immutable builders

[![GoDoc](https://godoc.org/github.com/lann/builder?status.png)](https://godoc.org/github.com/lann/builder)
[![Build Status](https://travis-ci.org/lann/builder.png?branch=master)](https://travis-ci.org/lann/builder)

Builder helps you write **fluent** interfaces for your libraries using method chaining:

```go
req := ReqBuilder.
    Url("http://golang.org").
    Header("User-Agent", "Builder").
    Get()
```

Builder uses **immutable** persistent data structures ([these](https://github.com/mndrix/ps), specifically)
so that each step in your method chain can be reused:

```go
build := WordBuilder.Letters("Build")
builder := build.Letters("er").String()
building := build.Letters("ing").String()
```

Builder makes it easy to **build** structs using the **builder** pattern (*surprise!*):

```go
muppet := builder.GetStruct(
    MuppetBuilder.
        Name("Beaker").
        Hair("orange"))

muppet.Name == "Beaker"
```

