"use strict"

var app = new Vue({
  el: "#app",
  components: {
    "home": httpVueLoader("web/vue/home.vue"),
  }, // --- End of components --- //

  data: {
    orders: [],
    commonClass: "btn p-3 border add-button-style",
  }, // --- End of data --- //

  created: function () {
    this.wsBase = "ws://localhost:4567/ws";
  }, // --- End of created --- //


  computed: {
  }, // --- End of computed --- //

  methods: {


  }, // --- End of methods --- //


  mounted: function () {
    let _this = this

    console.log("## mounted()");
    this.ws = new WebSocket(this.wsBase);
    this.ws.onopen = function (event) {
      _this.isOnline = true;
      console.log("### websocket.onopen()");
    };

    this.ws.onmessage = function (event) {
      console.log(event);




      const eventData = JSON.parse(event.data);
      let enableLog = true

      if (eventData.Key == "progress") {
        _this.methodVlcProgress(eventData.Value)
        enableLog = false
      }

      if (eventData.Key == "isDownloading") {
        _this.methodFileIsDownloading(eventData.Value)
      }

      if (eventData.Key == "files") {
        _this.methodFiles(eventData.Value)
      }

      if (eventData.Key == "devices") {
        _this.methodDevices(eventData.Value)
      }

      if (enableLog) {
        let message = {};
        message["📱＜ー💻"] = eventData.Room +" : "+ eventData.Object +" : "+ eventData.Key +" : "+ eventData.Value ;
        console.log(message);
      }
    };

    // websocketでエラーが発生した時
    this.ws.onerror = function (event) {
      console.log("### websocket.onerror()");
    };

    // websocketをクローズした時
    this.ws.onclose = function (event) {
      console.log("### websocket.onclose()");
      _this.isOnline = false;
      _this.timer = setInterval(function () {
        axios
          .get("")
          .then(function (response) {
            window.location.reload();
          })
          .catch(function (error) {
            console.log(error)
          })
      }, 1000);
    }
  }

})
