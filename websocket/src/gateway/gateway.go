package gateway

import (
	"proto"
	"webserver"
	"tcpserver"
	"dbproxy"
)

type OnClientActive 	func(uid uint32, addr string)
type OnClientClose 		func(uid uint32)
type OnClientMessage 	func(uint32 uint32, message *proto.Message)
type AuthClient 		func(uint32 uint32, addr string) error

type GateOption struct {
	Addr 			string
	Active 			OnClientActive
	Close 			OnClientClose
	Msg 			OnClientMessage
	Auth 			AuthClient
}

type Gateway struct {
	*ClientManager
	webServer 		*webserver.WebSocketServer
	tcpServer 		*tcpserver.TcpServer
	dbProxy 		*dbproxy.DbProxy
}

func NewGateway(opt *GateOption) *Gateway {
	gateway := &Gateway{
		ClientManager: NewClientManager(opt),
	}

	gateway.webServer = webserver.NewWebSocketServer(&webserver.WebOption{
		Addr: opt.Addr,
		Path: "/game",
		Activecb: func(client *webserver.WebClient) {
			gateway.OnClientConnected(client)
		},
		Closecb: func(client *webserver.WebClient) {
			gateway.OnClientDisconnct(client)
		},
		Authcb: func(client *webserver.WebClient) error {
			return gateway.OnClientAuth(client)
		},
		Msgcb: func(client *webserver.WebClient, message *proto.Message) {
			gateway.OnClientMessage(client, message)
		},
	})

	/*
	gateway.tcpServer = tcpserver.NewTcpServer(&tcpserver.TcpOption{
		Addr: ":9092",
		Activecb: func(cli *tcpserver.TcpClient) {
			fmt.Println("tcp connection connectd ")
		},
		Closecb: func(cli *tcpserver.TcpClient) {
			fmt.Println("tcp connection closed")
		},
		Auth: func(cli *tcpserver.TcpClient) error {
			fmt.Println("tcp connection authed")
			return nil
		},
		Msgcb: func(cli *tcpserver.TcpClient, m *proto.Message) {
			fmt.Println("tcp connection msgcb ", m)
		},
	})
	*/

	gateway.dbProxy = dbproxy.NewDbProxy(&dbproxy.DbOption{
		Host: "127.0.0.1:3306",
		User: "root",
		Pass: "1",
		Name: "mygame",
		ShowDetailLog: true,
		Singular: true,
	})

	return gateway
}

func (gw *Gateway) Start() {
	go gw.webServer.Start()
}


