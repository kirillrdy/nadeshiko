package main

import "github.com/kirillrdy/nadeshiko"

func main() {
	nadeshiko.Nadeshiko("/", handler)
	nadeshiko.Start()
}
