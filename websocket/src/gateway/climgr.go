package gateway

import (
	"webserver"
	"proto"
	"sync"
	"fmt"
)

type ClientManager struct {
	opt 			*GateOption
	idLock 			*sync.RWMutex
	clientIds 		map[uint32]*webserver.WebClient
	cliIdGen 		uint32

	gsOpt 			*GateOption
	lsOpt 			*GateOption

	lsMsgChanel 	*sync.RWMutex
	gsMsgChanel 	*sync.RWMutex
}

func NewClientManager(opt *GateOption) *ClientManager {
	return &ClientManager{
		opt: opt,
		idLock: new(sync.RWMutex),
		clientIds: make(map[uint32]*webserver.WebClient),
		cliIdGen: 0,
		lsMsgChanel: new(sync.RWMutex),
		gsMsgChanel: new(sync.RWMutex),
	}
}

func (cm *ClientManager) RegisterNoitfier(option *GateOption) {
	if option.Type == "gameserver" {
		cm.gsOpt = option
	} else if option.Type == "lobby" {
		cm.lsOpt = option
	}
}

func (cm *ClientManager) getClient(uid uint32) *webserver.WebClient {
	var client *webserver.WebClient
	cm.idLock.Lock()
	if c, ok := cm.clientIds[uid]; ok {
		client = c
	}
	cm.idLock.Unlock()
	return client
}

func (cm *ClientManager) OnClientConnected(client *webserver.WebClient) {
	var id uint32
	cm.idLock.Lock()
	cm.cliIdGen++
	id = cm.cliIdGen
	cm.clientIds[id] = client
	client.Uid = id
	cm.idLock.Unlock()

	//cm.opt.Active(id, client.ClientAddr())
	cm.lsOpt.Active(id, client.ClientAddr())
}

func (cm *ClientManager) OnClientDisconnct(client *webserver.WebClient) {
	//cm.opt.Close(client.Uid)
	cm.lsOpt.Close(client.Uid)
}

func (cm *ClientManager) OnClientAuth(client *webserver.WebClient) error {
	return nil
}

func (cm *ClientManager) OnClientMessage(client *webserver.WebClient, message *proto.Message) {
	if client != nil {
		//cm.msgChanel.Lock()
		//cm.opt.Msg(client.Uid, message)
		//cm.msgChanel.Unlock()
		if message.Cmd < 1000 {
			cm.lsMsgChanel.Lock()
			cm.lsOpt.Msg(client.Uid, message)
			cm.lsMsgChanel.Unlock()
		} else {
			cm.gsMsgChanel.Lock()
			cm.gsOpt.Msg(client.Uid, message)
			cm.gsMsgChanel.Unlock()
		}
	}
}

func (cm *ClientManager) KickClient(uid uint32) {
	client := cm.getClient(uid)
	if client != nil {
		client.ActiveClose()
	}
}

func (cm *ClientManager) SendMessage(uid uint32, message *proto.Message) {
	fmt.Println("gateway client send ", message)
	client := cm.getClient(uid)
	if client != nil {
		client.Send(message)
	}
}

func (cm *ClientManager) BroadcastMessage(ids []uint32, message *proto.Message) {
	for _, id := range ids {
		cm.SendMessage(id, message)
	}
}