
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
		
		addHandler: function(event, fn) {
			var handler = function(data) {	
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
			websocket.onclose = function(event) {
				self.onclose(event)
			}
	        websocket.onmessage = function(event) {
	        	self.onMessage(event)
	        }
	        websocket.onerror = function(event) {
	        	self.onerror(event)
	        }

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
			var data = JSON.parse(event.data)
			console.log("connection message " , data, this.handlers[data.cmd])
			var handler = this.handlers[data.cmd]
			if (handler) {
				console.log("connection process ", data.msg)
				handler(data.msg)
			} else {
				console.log("not have ", data.cmd);
			}
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
				console.log("send msg ", cmd, data)
			}
		},
	},
})
