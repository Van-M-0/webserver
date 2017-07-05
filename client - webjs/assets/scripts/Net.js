
if (window["WebSocket"] == null) {
	console.log("---------- warning cannot open websocket ------------")	
}

cc.Class({
	extends: cc.Component,

	statics: {
		handlers: {},
		ip: "127.0.0.1:9091",
		so: null,
		constatus: "closed",

		addHandler: function(evnet, fn) {
			var handler = function(data) {
				data = JSON.parse(data)
				fn(data)
			}
			this.handlers[event] = handler
		},

		removeHandler: function(event, fn) {
			if (this.handlers[event]) {
				this.handlers[event] = nil
				fn()
			}
		},

		connect: function(cb) {
			var self = this
			var websocket = new WebSocket("ws://"+this.ip+"/game")
			websocket.onopen = function(event) {
				self.onopen(event, cb)
			}
			websocket.onclose = this.onclose
	        websocket.onmessage = this.onMessage
	        websocket.onerror = this.onerror
	        this.so = websocket
	        this.constatus = "connecting"
		},

		onopen: function(event, cb) {
			console.log("connection open " + event)
			this.constatus = "connected"
			if (cb) {
				cb()
			}
		},

		onclose: function(event) {
			console.log("connection close " + event)
			this.constatus = "closed"
		},

		onMessage: function(event) {
			console.log("connection message " + event.data)
		},

		onerror: function(event) {
			console.log("connection error " + event)
			this.onclose()
		},

		send: function(cmd, data) {
			var msg = {
				cmd: cmd,
				msg: data
			}

			msg = JSON.stringify(msg)

			var self = this
			if (self.constatus == "closed") {
				self.connect(function() {
					var ret = self.so.send(msg)
					console.log("send msg ", cmd, data, ret)
				})
			} else {
				self.so.send(msg)
				console.log("send msg ", cmd, data, ret)
			}
		}
	},
})
