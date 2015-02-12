package nadeshiko

import "fmt"

//TODO consider depricating this
func (document *Document) SetLocalStorageItem(key, val string) {
	stringToSend := fmt.Sprintf("localStorage.setItem('%s','%s')", key, val)
	document.SendMessage(stringToSend)
}

//TODO consider depricating this
func (document *Document) GetLocalStorageItem(key string, callback func(string)) {
	randomString := generateCallbackId()
	callbacks[randomString] = overSocketCallback{oneTime: true, callback: func(vals ...string) {
		callback(vals[1])
	}}

	stringToSend := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",localStorage.getItem('%s')])); ", randomString, key)
	document.SendMessage(stringToSend)
}
