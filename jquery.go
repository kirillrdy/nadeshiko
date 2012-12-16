package nadeshiko

import "fmt"
import "strconv"
import "time"


type JQuerySelectedElements struct {
	selector string
	ws       *WebsocketConnection
}

func (ws *WebsocketConnection) JQuery(selector string) (element JQuerySelectedElements) {
	element.selector = selector
	element.ws = ws
	return
}


func (element JQuerySelectedElements) Append(content string) {
	element.oneArgumentMethod("append",content)
}


//TODO get rid of this method, and figure out more neat way of chaining jquery methods
func (element JQuerySelectedElements) PrevRemove() {
	string_to_send := fmt.Sprintf("$('%s').prev().remove()", element.selector)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) Before(content string) {
	element.oneArgumentMethod("before",content)
}

func (element JQuerySelectedElements) PrependString(content string) {
	element.oneArgumentMethod("prepend",content)
}

func (element JQuerySelectedElements) SetVal(new_value string) {
	element.oneArgumentMethod("val",new_value)
}

func (element JQuerySelectedElements) SetText(new_value string) {
	element.oneArgumentMethod("text",new_value)
}

func (element JQuerySelectedElements) Empty() {
	element.zeroArgumentMethod("empty")
}

func (element JQuerySelectedElements) Remove() {
	element.zeroArgumentMethod("remove")
}


func (element JQuerySelectedElements) Click(callback func()) {
	callback_id := generateCallbackId()

	Callbacks[callback_id] = OverSocketCallback{element.ws, false, func(...string) {
		callback()
	}}

	string_to_send := fmt.Sprintf("$('%s').click(function(){ ws.send(JSON.stringify([\"%x\"])); });", element.selector, callback_id)
	element.ws.SendMessage(string_to_send)
}

//TODO refactor function body, not DRY
func (element JQuerySelectedElements) Keydown(callback func(int)) {

	callback_id := generateCallbackId()

	Callbacks[callback_id] = OverSocketCallback{element.ws, false, func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	string_to_send := fmt.Sprintf("$('%s').keydown(function(event){ ws.send(JSON.stringify([\"%s\",event.keyCode.toString()])); });", element.selector, callback_id)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) GetVal(callback func(string)) {
	random_string := generateCallbackId()

	Callbacks[random_string] = OverSocketCallback{element.ws, true, func(vals ...string){
		callback(vals[1])
	}}

	string_to_send := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",$('%s').val()])); ", random_string, element.selector)
	element.ws.SendMessage(string_to_send)
}

////////////////////////////////////////////////////////////////////////////////////
// Unexported functions go here

func generateCallbackId() string {
	now := time.Now()
	//TODO get better way of generating uniq number
	random_number := now.UnixNano()
	return fmt.Sprintf("%x",random_number)
}

func (element JQuerySelectedElements) oneArgumentMethod(name string, param string) {
	string_content := strconv.Quote(param)
	string_to_send := fmt.Sprintf("$('%s').%s(%s)", element.selector, name, string_content)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) zeroArgumentMethod(name string) {
	string_to_send := fmt.Sprintf("$('%s').%s()", element.selector, name)
	element.ws.SendMessage(string_to_send)
}
