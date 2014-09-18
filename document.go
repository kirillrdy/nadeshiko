package nadeshiko

import (
	"log"
	"strings"

	"code.google.com/p/go.net/websocket"
)

type Document struct {
	websocket           *websocket.Conn
	current_transaction []string
	in_transaction      bool
	Error               error
	ClientDisconnected  bool
}

func (document *Document) StartBuffer() {
	document.in_transaction = true
}

func (document *Document) FlushBuffer() {
	document.in_transaction = false
	if len(document.current_transaction) != 0 {
		document.SendMessage(strings.Join(document.current_transaction, ";"))
		document.current_transaction = []string{}
	}
}

func (document *Document) SendMessage(message string) {

	if document.in_transaction {
		document.current_transaction = append(document.current_transaction, message)
	} else {
		err := websocket.Message.Send(document.websocket, message)

		if err != nil {
			log.Printf("ERROR '%s'\n", err)
			document.Error = err
			//runtime.Goexit()
		}
	}

	// This is done to cleanup timers that are not terminated
	// but will try to send on close sockets
	// perhaps we can do this when we change activities
	//TODO handle errors other than send on closed connection

	if Verbose {
		log.Printf("send: %s\n", message)
	}
}
