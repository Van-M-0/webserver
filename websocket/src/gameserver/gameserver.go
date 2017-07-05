package gameserver

import (
	"net"
	"fmt"
)

type GameServer struct {
	gwconn 		net.Conn
}

func NewGameServer() *GameServer {
	return &GameServer{

	}
}

func (gs *GameServer) Start() {
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		fmt.Println("game server conn errror", err)
		return
	}
	gs.gwconn = conn

}


