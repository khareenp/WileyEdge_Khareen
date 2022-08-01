// api/index.js
var socket = new WebSocket("ws://localhost:8080/ws");

// triggers a callback whenever it receives a new message from our WebSocket connection
let connect = (cb: (msg: string) => void) => {
  console.log("connecting");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = (msg: MessageEvent<string>) => {
    console.log(msg);
    cb(msg.data); //whenever message is received cb will be called
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};
let sendMsg = (msg: string) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
