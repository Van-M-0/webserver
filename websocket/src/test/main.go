package main

import (
	"sync"
	"scenemgr"
	"gameserver"
	"gateway"
	"lobby"
	"encoding/json"
	"fmt"
	"proto"
)

func test() {
	jstr :=  `{"account":"guest_123","conf":{"difen":123}}`
	var m proto.CreateRoomXZ_C
	err := json.Unmarshal([]byte(jstr), &m)
	fmt.Println(err, "----", m.Conf)
}

func main() {
	test()

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
