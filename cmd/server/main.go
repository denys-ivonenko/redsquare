package main

import (
	"github.com/denys-ivonenko/redsquare/internal"
)

func main() {
	e := internal.EchoRouter()
	e.Logger.Fatal(e.Start(":8181"))
}
