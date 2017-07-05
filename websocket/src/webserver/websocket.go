package webserver

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"proto"
)

type ActiveCallback func(client *WebClient)
type CloseCallback 	func(client *WebClient)
type MessageCallback func(client *WebClient, message *proto.Message)
type AuthCallback 	func(client *WebClient) error

type WebOption struct {
	Addr 		string
	Path 		string
	Activecb 	ActiveCallback
	Closecb 	CloseCallback
	Msgcb 		MessageCallback
	Authcb 		AuthCallback
}

type HttpOption struct {

}

type WebSocketServer struct {
	Opts 		*WebOption
	upgrader 	*websocket.Upgrader
}

func NewWebSocketServer(opts *WebOption) *WebSocketServer {
	return &WebSocketServer{
		Opts: opts,
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (ws *WebSocketServer) Start() {
	http.HandleFunc(ws.Opts.Path, ws.serve)
	if err := http.ListenAndServe(ws.Opts.Addr, nil); err != nil {
		fmt.Println("websocekt server start errror ", err)
	}
}

func (ws *WebSocketServer) serve(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade websocket connection err ", err)
		return
	}

	cli := NewWebClient(ws.Opts, conn)
	cli.start()
}

