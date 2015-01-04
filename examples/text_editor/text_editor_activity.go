package main

import (
	"fmt"
	"log"

	"github.com/kirillrdy/nadeshiko"
)

var clients []nadeshiko.Document

func handler(document nadeshiko.Document) {

	//go func() {
	//	blinkCursor(connection)
	//}()

	for i := 0; i < textBuffer.NumberOfLines(); i++ {
		document.JQuery("body").Append(fmt.Sprintf("<div id='%d'>%s</div>", i, textBuffer.Line(i)))
	}

	clients = append(clients, document)

	document.JQuery("body").Keydown(func(key int) {
		onKeyDown(document, key)
	})

	//document.JQuery("body").Keypress(func(key int) {
	//	activity.onKeyPress(document, key)
	//})
}

const BACK_SPACE = 8
const LEFT_KEY = 37
const UP_KEY = 38
const RIGHT_KEY = 39
const DOWN_KEY = 40
const ENTER_KEY = 13

func onKeyDown(document nadeshiko.Document, key int) {
	log.Printf("key down: %d \n", key)

	//if key == ENTER_KEY {
	//	y = y + 1
	//	addNewLine(connection, y)
	//	moveCursorToLine(connection, y)
	//} else if key == 8 {
	//	connection.JQuery("#cursor").PrevRemove()
	//} else if key == UP_KEY {
	//	y = y - 1
	//	moveCursorToLine(connection, y)
	//} else if key == DOWN_KEY {
	//	y = y + 1
	//	moveCursorToLine(connection, y)
	//}
}

func updateLine(line_number int) {

	for _, client := range clients {
		client.JQuery(fmt.Sprintf("#%d", line_number)).SetText(textBuffer.data[line_number])
	}
}
