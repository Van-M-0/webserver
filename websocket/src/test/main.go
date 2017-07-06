package main

import (
	"gateway"
	"sync"
)

func main() {
	gw := gateway.NewGateway()
	gw.Start()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Wait()
}
