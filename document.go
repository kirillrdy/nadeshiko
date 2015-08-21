package nadeshiko

import (
	"log"
	"strings"

	"golang.org/x/net/websocket"
)

//Document repesents clients document ( browser tab )
//is used for sending low level messages to the client
type Document struct {
	websocket          *websocket.Conn
	currentTransaction []string
	inTransaction      bool
	Error              error
	ClientDisconnected bool
}

//TODO not thread safe
func (document *Document) StartBuffer() {
	document.inTransaction = true
}

//TODO not thread safe
func (document *Document) FlushBuffer() {
	document.inTransaction = false
	if len(document.currentTransaction) != 0 {
		document.sendMessage(strings.Join(document.currentTransaction, ";"))
		document.currentTransaction = []string{}
	}
}

func (document *Document) sendMessage(message string) {

	if document.inTransaction {
		document.currentTransaction = append(document.currentTransaction, message)
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

	if verbose {
		log.Printf("send: %s\n", message)
	}
}
