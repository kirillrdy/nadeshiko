package nadeshiko

import "fmt"

func SocketJSCode(path string) string {
	//TODO escape path
	//TODO rewrite in something that compiles to js
	template := `
function init(reconnecting) {
  var host = window.location.hostname
  var port = window.location.port
  var path = window.location.pathname
  var web_socket_location = "ws://"+host+":" + port + "%v"
  ws = new WebSocket(web_socket_location);

  ws.onopen = function(){
    reconnecting = false
  }

  ws.onmessage = function (e) {
     eval(e.data);

     //TODO write to log only in dev env
     console.log(e.data)
  };
  
  ws.onclose = function(){
    if (! reconnecting) {
      $(document.body).append("<span  class=\"label label-warning connection-lost\">Lost connection retrying</span>")
    }
    setTimeout("init(true)",500)
  };
};

$(function(){
  init(false);
})
`
	return fmt.Sprintf(template, path)
}
