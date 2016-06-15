var s = new WebSocket("ws://localhost:8080/ws");

s.onmessage = function(event) {
  console.log(event.data);
};
