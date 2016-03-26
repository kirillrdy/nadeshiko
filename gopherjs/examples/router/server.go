// +build !js

package main

import (
	"github.com/kirillrdy/nadeshiko/gopherjs"
	"log"
	"net/http"
)

func main() {

	app := gopherjs.App{PackageName: "github.com/kirillrdy/nadeshiko/gopherjs/examples/router"}
	app.Mount("/")

	address := ":3000"
	log.Printf("Starting server on %s", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Panic(err)
	}
}
