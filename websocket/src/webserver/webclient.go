package webserver

import (
	"github.com/gorilla/websocket"
	"proto"
	"github.com/ugorji/go/codec"
	"fmt"
)

type WebClient struct {
	conn 		*websocket.Conn
	Opts 		*WebOption
	sendch 		chan *proto.Message
	sendbuf 	[]byte
}

func NewWebClient(opt *WebOption, conn *websocket.Conn) *WebClient {
	return &WebClient{
		conn: conn,
		Opts: opt,
		sendch: make(chan *proto.Message),
		sendbuf: make([]byte, 10240),
	}
}

func (wb *WebClient) start() {
	wb.Opts.Activecb(wb)
	if wb.Opts.Authcb != nil {
		if err := wb.Opts.Authcb(wb); err != nil {
			return
		}
	}
	go wb.readLoop()
	go wb.writeLoop()
}

func (wb *WebClient) close() {
	wb.Opts.Closecb(wb)
	wb.conn.Close()
}

func (wb *WebClient) readLoop() {
	defer func() {
		wb.close()
	}()

	decodec := new(codec.JsonHandle)
	for {
		mt, raw, err := wb.conn.ReadMessage()
		if err != nil {
			fmt.Println("web socket conn recv msg err ", err)
			return
		}

		fmt.Println("web conn recv message ", mt, raw, err)

		var msg proto.Message
		if err := codec.NewDecoderBytes(raw, decodec).Decode(&msg); err != nil {
			fmt.Println("web socket conn decode msg err ", err)
			return
		}

		wb.Opts.Msgcb(wb, &msg)
	}
}

func (wb *WebClient) writeLoop() {
	for {
		select {
		case m := <- wb.sendch:
			wb.sendEncode(m)
		}
	}
}

func (wb *WebClient) Send(m *proto.Message) {
	wb.sendch <- m
}

func (wb *WebClient) sendEncode(m *proto.Message) {
	encoder := codec.NewEncoderBytes(&wb.sendbuf, new(codec.JsonHandle))
	if err := encoder.Encode(m); err != nil {
		fmt.Println("encode msg error", err)
	}

	fmt.Println("write message ", wb.sendbuf, m)
	w, err := wb.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	w.Write(wb.sendbuf)

	if err := w.Close(); err != nil {
		return
	}
}