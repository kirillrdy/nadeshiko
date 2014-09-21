function init(handler, reconnecting) {
  var host = window.location.hostname
  var port = window.location.port
  var web_socket_location = "ws://"+host+":" + port + "/" + handler
  ws = new WebSocket(web_socket_location);

  ws.onopen = function(){
    $(document.body).empty()
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
    setTimeout("init(\".websocket\", true)",500)
  };
};

$(function(){
  init(".websocket",false);
})
