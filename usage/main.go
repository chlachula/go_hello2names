package main

import (
	good "github.com/chlachula/go_hello2names"
	"github.com/chlachula/go_hello2names/a"
	"github.com/chlachula/go_hello2names/b"
)

func greetings() string {
	return "Hello " + a.NameA() + " & " + b.NameB()
}

func text() string {
	return greetings() + ". " + good.Message()
}

func main() {
	print(text())
}
