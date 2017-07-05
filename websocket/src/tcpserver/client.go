package tcpserver

import (
	"net"
	"proto"
	"fmt"
	"github.com/ugorji/go/codec"
	"io"
)

type TcpClient struct {
	conn 	net.Conn
	header 	[2]byte
	body 	[]byte
	sendch	chan *proto.Message
	sendbuf []byte
}

func (cli *TcpClient) Close() {

}

func (cli *TcpClient) Start() {

}

func (cli *TcpClient) sendEncode(m *proto.Message) {
	encoder := codec.NewEncoderBytes(&cli.sendbuf, new(codec.MsgpackHandle))
	if err := encoder.Encode(m); err != nil {
		fmt.Println("encode msg error", err)
		return
	}
	cli.conn.Write(cli.sendbuf)
}

func (cli *TcpClient) Send(m *proto.Message) {
	cli.sendch <- m
}

func (cli *TcpClient) SendLoop() {
	for {
		select {
		case m := <- cli.sendch:
			cli.sendEncode(m)
		default:
			fmt.Println("send chan default")
		}
	}
}

func NewDailClient(opt *TcpOption) (*TcpClient, error) {
	conn, err := net.Dial("tcp", opt.Addr)
	if err != nil {
		return nil, err
	}

	cli := &TcpClient{
		conn: conn,
		body: make([]byte, 4096),
		sendch: make(chan *proto.Message),
		sendbuf: make([]byte, 8192),
	}

	opt.Activecb(cli)

	if opt.Auth {
		if err := opt.Auth(cli); err != nil {
			return nil, err
		}
	}
	go cli.SendLoop()

	recv := func() {

		defer func() {
			cli.Close()
			opt.Closecb(cli)
		}()

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

			opt.Msgcb(cli, &msg)
		}
	}

	go recv()

	return cli, nil
}
