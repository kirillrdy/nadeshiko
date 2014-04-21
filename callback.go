package nadeshiko

type overSocketCallback struct {
	connection *Connection
	oneTime    bool
	callback   func(...string)
}

var callbacks = make(map[string]overSocketCallback)
var deleteCallback = make(chan string)

func init() {
	go func() {
		for callbackId := range deleteCallback {
			delete(callbacks, callbackId)
		}
	}()
}
