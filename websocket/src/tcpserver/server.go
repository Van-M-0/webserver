package tcpserver

import (
	"proto"
	"net"
	"io"
	"fmt"
	"github.com/ugorji/go/codec"
)

type ActiveCallback func(cli *TcpClient)
type CloseCallback func(cli *TcpClient)
type MessageCallback func(cli *TcpClient, m *proto.Message)
type AuthClient	func(cli *TcpClient) error

type TcpOption struct {
	Addr 			string
	Activecb 		ActiveCallback
	Closecb 		CloseCallback
	Msgcb 			MessageCallback
	Auth 			AuthClient
}

type TcpServer struct {
	Opts 		*TcpOption
}

func NewTcpServer(opts *TcpOption) *TcpServer {
	return &TcpServer{
		Opts: opts,
	}
}

func (server *TcpServer) Start() error {
	l, err := net.Listen("tcp", server.Opts.Addr)
	if err != nil {
		return err
	}

	defer func() {
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}

		go server.ServeClient(conn)
	}
}

func (server *TcpServer) ServeClient(conn net.Conn) {

	cli := &TcpClient{
		conn:conn,
		body: make([]byte, 4096),
		sendch: make(chan *proto.Message),
		sendbuf: make([]byte, 8192),
	}
	defer func() {
		cli.Close()
		server.Opts.Closecb(cli)
	}()

	server.Opts.Activecb(cli)

	if server.Opts.Auth {
		if err := server.Opts.Auth(cli); err != nil {
			return
		}
	}
	go cli.SendLoop()

	decodec := new(codec.MsgpackHandle)
	for {
		size, err := io.ReadFull(conn, cli.header[:])
		if err != nil {
			break
		}

		var data []byte
		if size > len(cli.body) {
			fmt.Println("recv msg over max :", size)
			data = make([]byte, size)
		} else {
			data = cli.body
		}

		var msg proto.Message
		if err := codec.NewDecoderBytes(data, decodec).Decode(&msg); err != nil {
			break
		}

		server.Opts.Msgcb(cli, &msg)
	}
}

func (server *TcpServer) Stop() {

}


