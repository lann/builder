# Builder - fluent immutable builders for Go

[![GoDoc](https://godoc.org/github.com/lann/builder?status.png)](https://godoc.org/github.com/lann/builder)
[![Build Status](https://travis-ci.org/lann/builder.png?branch=master)](https://travis-ci.org/lann/builder)

Builder helps you write **fluent** DSLs for your libraries with method chaining:

```go
resp := ReqBuilder.
    Url("http://golang.org").
    Header("User-Agent", "Builder").
    Get()
```

Builder uses **immutable** persistent data structures ([these](https://github.com/mndrix/ps), specifically)
so that each step in your method chain can be reused:

```go
build := WordBuilder.AddLetters("Build")
builder := build.AddLetters("er")
building := build.AddLetters("ing")
```

Builder makes it easy to **build** structs using the **builder** pattern (*surprise!*):

```go
import "github.com/lann/builder"

type Muppet struct {
    Name string
    Friends []string
}

type muppetBuilder builder.Builder

func (b muppetBuilder) Name(name string) muppetBuilder {
    return builder.Set(b, "Name", name).(muppetBuilder)
}

func (b muppetBuilder) AddFriend(friend string) muppetBuilder {
    return builder.Append(b, "Friends", friend).(muppetBuilder)
}

var MuppetBuilder = builder.Register(muppetBuilder{}, Muppet{}).(muppetBuilder)
```
```go
builder.GetStruct(
    MuppetBuilder.
        Name("Beaker").
        AddFriend("Dr. Honeydew"))

Muppet{Name:"Beaker", Friends:[]string{"Dr. Honeydew"}}
```

