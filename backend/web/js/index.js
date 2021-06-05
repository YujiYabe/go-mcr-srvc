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
    methodChangeRoom: function (value) {
      if (this.room != value) {
        this.room = value;
        if (this.room == "bed") {
          this.device = this.deviceBed
          this.airconStatus = this.airconBedStatus
          this.airconWarm = this.airconBedWarm
          this.airconCool = this.airconBedCool
          this.lightStatus = this.lightBedStatus

        }
        if (this.room == "living") {
          this.device = this.deviceLiving
          this.airconStatus = this.airconLivingStatus
          this.airconWarm = this.airconLivingWarm
          this.airconCool = this.airconLivingCool
          this.lightStatus = this.lightLivingStatus
        }
        this.methodVibrate()
        localStorage.setItem("room", this.room)
        localStorage.setItem("device", this.device)
      }
    },

    methodChangeDevice: function (value) {
      this.device = value;
      if (this.room == "bed") {
        this.deviceBed = value;
      }
      if (this.room == "living") {
        this.deviceLiving = value;
      }
      this.methodVibrate()
      localStorage.setItem("room", this.room)
      localStorage.setItem("device", this.device)
    },

    methodSendCommand: function (value) {
      let args = {};
      args["Room"] = value.Room == undefined ? this.room : value.Room;
      args["Object"] = value.Object == undefined ? this.device : value.Object;
      args["Key"] = value.Key;
      args["Value"] = value.Value;

      if (args["Room"] == null || args["Object"] == null || args["Key"] == null || args["Value"] == null) {
        console.log(args)
      } else {
        this.methodVibrate(value.VibrateDisable)
        this.ws.send(JSON.stringify(args));
        let message = {};
        message["📱ー＞💻"] = args["Room"] + " : " + args["Object"] + " : " + args["Key"] + " : " + args["Value"] ;
        console.log(message);
      }
    },

    methodVlcProgress: function (value) {
      this.vlcProgress = parseFloat(value);
    },

    methodFileIsDownloading: function (value) {
      if (value == "true"){
        this.isDownloading = true;
      }else{
        this.isDownloading = false;
      }
    },

    methodFiles: function (json_string) {
      const value = JSON.parse(json_string);
      this.files = value;
    },

    methodFileUpload: function (value) {
      let formData = new FormData();
      formData.append("file", value);
      let config = {
        headers: {
          "content-type": "multipart/form-data"
        }
      };

      this.isUploading = true
      let _this = this

      axios
        .post("file_upload", formData, config)
        .then(function (response) {
          _this.isUploading = false
          console.log(response)
        })
        .catch(function (error) {
          _this.isUploading = false
          console.log(error)
        })
    },

    methodDevices: function (json_string) {
      //light setting
      const value = JSON.parse(json_string);
      this.lightBedStatus = value["bed"]["Light"]["Status"];
      this.lightLivingStatus = value["living"]["Light"]["Status"];
      if (this.room == "bed") {
        this.lightStatus = this.lightBedStatus;
      }
      if (this.room == "living") {
        this.lightStatus = this.lightLivingStatus;
      }

      //aircon setting
      this.airconBedStatus = value["bed"]["Aircon"]["Status"];
      this.airconBedWarm = value["bed"]["Aircon"]["WarmTemperature"];
      this.airconBedCool = value["bed"]["Aircon"]["CoolTemperature"];
      this.airconLivingStatus = value["living"]["Aircon"]["Status"];
      this.airconLivingWarm = value["living"]["Aircon"]["WarmTemperature"];
      this.airconLivingCool = value["living"]["Aircon"]["CoolTemperature"];
      if (this.room == "bed") {
        this.airconStatus = this.airconBedStatus
        this.airconWarm = this.airconBedWarm
        this.airconCool = this.airconBedCool
      }
      if (this.room == "living") {
        this.airconStatus = this.airconLivingStatus
        this.airconWarm = this.airconLivingWarm
        this.airconCool = this.airconLivingCool
      }
    },

    methodVibrate (isDisable) {
      if (!isDisable) {
        window.navigator.vibrate(10);
      }
    },
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
