package main

import (
	"time"

	"github.com/kirillrdy/nadeshiko"
)

type Cursor struct {
	x, y int
}

func blinkCursor(document nadeshiko.Document) {
	tick := time.Tick(500 * time.Millisecond)
	show_cursor := false
	for _ = range tick {
		show_cursor = !show_cursor
		if show_cursor {
			document.JQuery("#cursor").SetText("|")
		} else {
			document.JQuery("#cursor").SetText("")
		}
	}
}
