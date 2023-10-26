
function setupWebSocket () {
  const kind = $('#kind').attr('name');
  const number = $('#number').attr('name');

  const urlString = "wss://" + document.domain + ":" + location.port + "/v1/ws/app?kind=" + kind + "&number=" + number;
  let socket = new WebSocket(urlString);

  socket.onopen = function (event) {
    console.log("Connected to the WebSocket server");
  };

  socket.onmessage = function (event) {
    buildOnWs(event.data);
  };

  socket.onerror = function (error) {
    console.log(`Error occurred: ${error.message}`);
  };

  socket.onclose = function (event) {
    if (event.wasClean) {
      console.log(`Connection closed cleanly, code=${event.code}, reason=${event.reason}`);
    } else {
      console.log('Connection closed due to an error');
    }
    setTimeout(setupWebSocket, 1000); //  自動再接続 
  };
}

// 初期の接続
setupWebSocket();
