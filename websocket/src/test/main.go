package main

import (
	"sync"
	"scenemgr"
	"gameserver"
	"gateway"
	"lobby"
	"encoding/json"
	"fmt"
)

func test() {
	type MyStruct struct {
		Name 		string 		`json:"name"`
	}

	m := &MyStruct{
		Name: "hello",
	}

	var i interface{}
	i = m
	s := "{'name':'hello'}"
	b1 := []byte(s)

	json.Unmarshal(b1, i)
	m2, err := i.(*MyStruct)
	fmt.Println(m2.Name, err)
}

func main() {
	gw := gateway.NewGateway(&gateway.GateOption{
		Addr: ":9091",
	})

	sm := scenemgr.NewRoomSceneManager()
	gs := gameserver.NewSingleGameServer(sm, gw)
	gs.Start()

	lm := scenemgr.NewLobbyceneManager()
	ls := lobby.NewSingleLobbyServer(lm, gw)
	ls.Start()

	gw.Start()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Wait()
}
