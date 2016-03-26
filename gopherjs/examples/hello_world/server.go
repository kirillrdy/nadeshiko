// +build !js

package main

import (
	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/gopherjs"
)

func main() {
	app := gopherjs.App{PackageName: "github.com/kirillrdy/nadeshiko/gopherjs/examples/hello_world"}
	app.Mount("/")

	nadeshiko.StartServer()
}
