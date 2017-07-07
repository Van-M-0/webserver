package main

import (
	"sync"
	"scenemgr"
	"gameserver"
)

func main() {

	sm := scenemgr.NewSceneManager()
	gs := gameserver.NewSingleGameServer(sm, ":9091")
	gs.Start()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Wait()
}
