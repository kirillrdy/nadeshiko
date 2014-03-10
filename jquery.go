package nadeshiko

import (
	"fmt"
	"strconv"
	"time"
	"io/ioutil"
)

type JQuerySelectedElements struct {
	selector   string
	connection *Connection
}

func (connection *Connection) JQuery(selector string) (element JQuerySelectedElements) {
	element.selector = selector
	element.connection = connection
	return
}

func (element JQuerySelectedElements) Write(content []byte) (int, error) {
	element.Append(string(content))
	return len(content), nil
}

func (element JQuerySelectedElements) Append(content string) {
	element.oneArgumentMethod("append", content)
}

//TODO get rid of this method, and figure out more neat way of chaining jquery methods
func (element JQuerySelectedElements) PrevRemove() {
	string_to_send := fmt.Sprintf("$('%s').prev().remove()", element.selector)
	element.connection.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) Before(content string) {
	element.oneArgumentMethod("before", content)
}

func (element JQuerySelectedElements) PrependString(content string) {
	element.oneArgumentMethod("prepend", content)
}

func (element JQuerySelectedElements) SetVal(new_value string) {
	element.oneArgumentMethod("val", new_value)
}

func (element JQuerySelectedElements) SetAttr(attr, new_value string) {
	element.twoArgumentMethod("attr", attr, new_value)
}

func (element JQuerySelectedElements) SetText(new_value string) {
	element.oneArgumentMethod("text", new_value)
}

func (element JQuerySelectedElements) Empty() {
	element.zeroArgumentMethod("empty")
}

func (element JQuerySelectedElements) Remove() {
	element.zeroArgumentMethod("remove")
}

func (element JQuerySelectedElements) Click(callback func()) {
	element.zeroArgumentMethodWithCallback("click", callback)
}

func (element JQuerySelectedElements) Change(callback func()) {
	element.zeroArgumentMethodWithCallback("change", callback)
}

//TODO refactor function body, not DRY
func (element JQuerySelectedElements) Keydown(callback func(int)) {

	callback_id := generateCallbackId()

	callbacks[callback_id] = overSocketCallback{element.connection, false, func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	string_to_send := fmt.Sprintf("$('%s').keydown(function(event){ ws.send(JSON.stringify([\"%s\",event.keyCode.toString()])); });", element.selector, callback_id)
	element.connection.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) GetVal(callback func(string)) {
	random_string := generateCallbackId()

	callbacks[random_string] = overSocketCallback{element.connection, true, func(vals ...string) {
		callback(vals[1])
	}}

	string_to_send := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",$('%s').val()])); ", random_string, element.selector)
	element.connection.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) LoadHtmlFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil { panic(err)}
	element.Append(string(data))
}
////////////////////////////////////////////////////////////////////////////////////
// Unexported functions go here

func generateCallbackId() string {
	now := time.Now()
	//TODO get better way of generating uniq number
	random_number := now.UnixNano()
	return fmt.Sprintf("%x", random_number)
}

func (element JQuerySelectedElements) twoArgumentMethod(name,  param1, param2 string) {
	quoted1 := strconv.Quote(param1)
	quoted2 := strconv.Quote(param2)
	string_to_send := fmt.Sprintf("$('%s').%s(%s,%s)", element.selector, name, quoted1, quoted2)
	element.connection.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) oneArgumentMethod(name string, param string) {
	string_content := strconv.Quote(param)
	string_to_send := fmt.Sprintf("$('%s').%s(%s)", element.selector, name, string_content)
	element.connection.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) zeroArgumentMethod(name string) {
	string_to_send := fmt.Sprintf("$('%s').%s()", element.selector, name)
	element.connection.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) zeroArgumentMethodWithCallback(name string, callback func()) {
	callback_id := generateCallbackId()

	callbacks[callback_id] = overSocketCallback{element.connection, false, func(...string) {
		callback()
	}}

	string_to_send := fmt.Sprintf("$('%s').%s(function(){ ws.send(JSON.stringify([\"%s\"])); });", element.selector, name, callback_id)
	element.connection.SendMessage(string_to_send)
}
