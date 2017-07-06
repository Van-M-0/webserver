package gateway

import (
	"proto"
	"webserver"
	"fmt"
	"tcpserver"
	"dbproxy"
)

type Gateway struct {
	webServer 		*webserver.WebSocketServer
	tcpServer 		*tcpserver.TcpServer
	dbProxy 		*dbproxy.DbProxy
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
	gw.InitDabase()

	go gw.webServer.Start()
	//go gw.tcpServer.Start()
}

func (gw *Gateway) InitDabase() {
	gw.dbProxy.CreateTableIfNot(&proto.T_Accounts{})
	gw.dbProxy.CreateTableIfNot(&proto.T_Games{})
	gw.dbProxy.CreateTableIfNot(&proto.T_GamesArchive{})
	gw.dbProxy.CreateTableIfNot(&proto.T_Guests{})
	gw.dbProxy.CreateTableIfNot(&proto.T_Message{})
	gw.dbProxy.CreateTableIfNot(&proto.T_Rooms{})
	gw.dbProxy.CreateTableIfNot(&proto.T_RoomUser{})
}

