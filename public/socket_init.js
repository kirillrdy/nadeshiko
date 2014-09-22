function init(reconnecting) {
  var host = window.location.hostname
  var port = window.location.port
  var path = window.location.pathname
  var web_socket_location = "ws://"+host+":" + port + path + ".websocket"
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
    setTimeout("init(true)",500)
  };
};

$(function(){
  init(false);
})
