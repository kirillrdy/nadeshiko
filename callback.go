package nadeshiko

import "sync"

type overSocketCallback struct {
	connection *Connection
	oneTime    bool
	callback   func(...string)
}

var callbacks = make(map[string]overSocketCallback)
var callbacksMutex sync.Mutex

func deleteCallback(callbackId string) {
	callbacksMutex.Lock()
	delete(callbacks, callbackId)
	callbacksMutex.Unlock()
}
