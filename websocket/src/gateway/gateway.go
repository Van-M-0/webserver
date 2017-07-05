package gateway

import (
	"proto"
	"webserver"
	"fmt"
)

type Gateway struct {
	webServer 		*webserver.WebSocketServer
}

func NewGateway() *Gateway {
	gateway := &Gateway{}
	gateway.webServer = webserver.NewWebSocketServer(&webserver.WebOption{
		Addr: ":9091",
		Path: "/game",
		Activecb: func(client *webserver.WebClient) {
			fmt.Println("web connection connected")
		},
		Closecb: func(client *webserver.WebClient) {
			fmt.Println("web connection closed")
		},
		Authcb: func(client *webserver.WebClient) error {
			return nil
		},
		Msgcb: func(client *webserver.WebClient, message *proto.Message) {
			fmt.Println("web connection message", message)
			if message.Cmd == 101 {
				client.Send(&proto.Message{
					Cmd: 102,
					Msg: &proto.GuestAuthMessage{
						Account: "hello world",
					},
				})
			}
		},
	})

	return gateway
}

func (gw *Gateway) Start() {
	gw.webServer.Start()
}


