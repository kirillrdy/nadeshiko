package nadeshiko

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

type jQuerySelector struct {
	//TODO change this from string to css.Selector
	selector string
	document *Document
}

//JQuery allows calling JQuery functions on given css.Selector
func (document *Document) JQuery(selector css.Selector) (element jQuerySelector) {
	element.selector = selector.Selector()
	element.document = document
	return
}

//func (element jQuerySelector) Write(content []byte) (int, error) {
//	element.Append(string(content))
//	return len(content), nil
//}

// HTML
func (element jQuerySelector) HTML(content html.Node) {
	element.oneArgumentMethod("html", content.String())
}

func (element jQuerySelector) Append(content html.Node) {
	element.oneArgumentMethod("append", content.String())
}

func (element jQuerySelector) Before(content html.Node) {
	element.oneArgumentMethod("before", content.String())
}

func (element jQuerySelector) Prepend(content html.Node) {
	element.oneArgumentMethod("prepend", content.String())
}

func (element jQuerySelector) PrependString(content string) {
	element.oneArgumentMethod("prepend", content)
}

func (element jQuerySelector) SetVal(newValue string) {
	element.oneArgumentMethod("val", newValue)
}

func (element jQuerySelector) SetCSS(attr, newValue string) {
	element.twoArgumentMethod("css", attr, newValue)
}

func (element jQuerySelector) SetAttr(attr, newValue string) {
	element.twoArgumentMethod("attr", attr, newValue)
}

func (element jQuerySelector) SetText(newValue string) {
	element.oneArgumentMethod("text", newValue)
}

func (element jQuerySelector) Empty() {
	element.zeroArgumentMethod("empty")
}

func (element jQuerySelector) Remove() {
	element.zeroArgumentMethod("remove")
}

func (element jQuerySelector) Click(callback func()) {
	element.zeroArgumentMethodWithCallback("click", callback)
}

func (element jQuerySelector) On(eventName string, callback func()) {
	element.oneArgumentMethodWithCallback("on", eventName, callback)
}

func (element jQuerySelector) Change(callback func()) {
	element.zeroArgumentMethodWithCallback("change", callback)
}

//TODO refactor function body, not DRY
func (element jQuerySelector) Keypress(callback func(int)) {

	callbackID := generateCallbackID()

	//TODO not threadsafe
	callbacks[callbackID] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	stringToSend := fmt.Sprintf("$('%s').keypress(function(event){ ws.send(JSON.stringify([\"%s\",event.charCode.toString()])); });", element.selector, callbackID)
	element.document.sendMessage(stringToSend)
}

//TODO refactor function body, not DRY
func (element jQuerySelector) Keydown(callback func(int)) {

	callbackID := generateCallbackID()

	//TODO not thread safe
	callbacks[callbackID] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	stringToSend := fmt.Sprintf("$('%s').keydown(function(event){ ws.send(JSON.stringify([\"%s\",event.keyCode.toString()])); });", element.selector, callbackID)
	element.document.sendMessage(stringToSend)
}

func (element jQuerySelector) GetVal(callback func(string)) {
	randomString := generateCallbackID()

	callbacks[randomString] = overSocketCallback{connection: element.document.websocket, oneTime: true, callback: func(vals ...string) {
		callback(vals[1])
	}}

	stringToSend := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",$('%s').val()])); ", randomString, element.selector)
	element.document.sendMessage(stringToSend)
}

func (element jQuerySelector) GetValChan() chan string {
	randomString := generateCallbackID()

	result := make(chan string)

	callback := func(vals ...string) {
		result <- vals[1]
	}

	callbacks[randomString] = overSocketCallback{connection: element.document.websocket, oneTime: true, callback: callback}

	stringToSend := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",$('%s').val()])); ", randomString, element.selector)
	element.document.sendMessage(stringToSend)
	return result
}

////////////////////////////////////////////////////////////////////////////////////
// Unexported functions go here

func generateCallbackID() string {
	now := time.Now()
	//TODO get better way of generating uniq number
	randomNumber := now.UnixNano()
	return fmt.Sprintf("%x", randomNumber)
}

func (element jQuerySelector) twoArgumentMethod(name, param1, param2 string) {
	quoted1 := strconv.Quote(param1)
	quoted2 := strconv.Quote(param2)
	stringToSend := fmt.Sprintf("$('%s').%s(%s,%s)", element.selector, name, quoted1, quoted2)
	element.document.sendMessage(stringToSend)
}

func (element jQuerySelector) oneArgumentMethod(name string, param string) {
	stringContent := strconv.Quote(param)
	stringToSend := fmt.Sprintf("$('%s').%s(%s)", element.selector, name, stringContent)
	element.document.sendMessage(stringToSend)
}

func (element jQuerySelector) zeroArgumentMethod(name string) {
	stringToSend := fmt.Sprintf("$('%s').%s()", element.selector, name)
	element.document.sendMessage(stringToSend)
}

func (element jQuerySelector) oneArgumentMethodWithCallback(name string, arg1 string, callback func()) {
	callbackID := generateCallbackID()

	callbacks[callbackID] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(...string) {
		callback()
	}}

	stringToSend := fmt.Sprintf("$('%s').%s('%s', function(){ ws.send(JSON.stringify([\"%s\"])); });", element.selector, name, arg1, callbackID)
	element.document.sendMessage(stringToSend)
}

func (element jQuerySelector) zeroArgumentMethodWithCallback(name string, callback func()) {
	callbackID := generateCallbackID()

	callbacks[callbackID] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(...string) {
		callback()
	}}

	stringToSend := fmt.Sprintf("$('%s').%s(function(){ ws.send(JSON.stringify([\"%s\"])); });", element.selector, name, callbackID)
	element.document.sendMessage(stringToSend)
}
