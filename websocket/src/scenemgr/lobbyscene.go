package scenemgr

import (
	"proto"
	"fmt"
	"dbproxy"
	"math/rand"
)

type LobbySceneMgr struct {
	manager 		GameManager
	dbProxy 		*dbproxy.DbProxy
}

func NewLobbyceneManager() *LobbySceneMgr {
	return &LobbySceneMgr{
		dbProxy: dbproxy.NewDbProxy(&dbproxy.DbOption{
			Host: "127.0.0.1:3306",
			User: "root",
			Pass: "1",
			Name: "mygame",
			ShowDetailLog: true,
			Singular: true,
		}),
	}
}

func (sm *LobbySceneMgr) RegisterScene(opt *SceneOption) error {
	return nil
}

func (sm *LobbySceneMgr) KickClient(id uint32, reason *proto.Message) {
	sm.manager.KickClient(id, reason)
}

func (sm *LobbySceneMgr) SendClientMessage(id uint32, m *proto.Message) {
	sm.manager.SendClientMessage(id, m)
}

func (sm *LobbySceneMgr) BroadcastMessage(ids []uint32, m *proto.Message) {
	sm.manager.BroadcastMessage(ids, m)
}

func (sm *LobbySceneMgr) SendCommandInfo(source uint32, m *proto.Message) {
	sm.manager.SendCommandInfo(source, m)
}

func (sm *LobbySceneMgr) OnInit(mgr GameManager) {
	sm.manager = mgr
}

func (sm *LobbySceneMgr) OnClientConnected(id uint32, addr string) {
	fmt.Println("lobby client connected ", id, addr)
}

func (sm *LobbySceneMgr) OnClientDisconnected(id uint32) {
	fmt.Println("lobby client didconnected ", id)
}

func (sm *LobbySceneMgr) OnClientMessage(id uint32, m *proto.Message) {
	fmt.Println("lobby scene manager recv message .. ", id, m)
	switch m.Cmd {
	case proto.CmdGuestAuth:
		sm.HandleGuest(id, m.Msg.(*proto.AuthGuest_C))
	case proto.CmdLoginLobby:
		sm.HandleLogin(id, m.Msg.(*proto.Login_C))
	case proto.CmdCreateAccount:
		sm.HandleCreateAccount(id, m.Msg.(*proto.CreateAccount_C))
	default:
		fmt.Println("recv unkown cmd ", m.Cmd)
	}
}

func (sm *LobbySceneMgr) OnCommandInfo(source uint32, m *proto.Message) {

}

func (sm *LobbySceneMgr) CreateScene(opt *SceneOption) Scene {
	return nil
}

func (sm *LobbySceneMgr) SendUserMessage(id uint32, cmd uint16, i interface{}) {
	sm.SendClientMessage(id, &proto.Message{
		Cmd: cmd,
		Msg: i,
	})
}

func (sm *LobbySceneMgr) HandleGuest(id uint32, c *proto.AuthGuest_C) {
	sm.SendUserMessage(id, proto.CmdGuestAuth, &proto.GuestAuthMessage {
		ErrCode: 0,
		Account: c.Account,
		Sign: "1234567890",
	})
}

func (sm *LobbySceneMgr) HandleLogin(id uint32, c *proto.Login_C) {
	acc := "guest_" + c.Account
	var userInfo proto.T_Users
	if sm.dbProxy.GetUserInfo(acc, &userInfo) {
		fmt.Println("handle login account already exists ", acc, c)
		sm.SendUserMessage(id, proto.CmdLoginLobby, &proto.LoginLobbyScucess{
			ErrCode: 0,
			Account: userInfo.Account,
			UserId: userInfo.Userid,
			Name: userInfo.Name,
			Level: userInfo.Level,
			Exp: userInfo.Exp,
			Coins: userInfo.Coins,
			Gems: userInfo.Gems,
			Sex: userInfo.Sex,
			RoomId: userInfo.Roomid,
			Ip: "123.123.123.132",
		})
	} else {
		fmt.Println("handle login account not exists ", acc, c)
		sm.SendUserMessage(id, proto.CmdLoginLobby, &proto.LoginLobbyScucess{
			ErrCode: 0,
		})
	}
}

func (sm *LobbySceneMgr) HandleCreateAccount(id uint32, c *proto.CreateAccount_C) {
	acc := "guest_" + c.Account
	var userInfo proto.T_Users
	fmt.Println("create account ", acc, c)
	if sm.dbProxy.GetUserInfo(acc, &userInfo) {
		fmt.Println("create account get", userInfo)
		sm.SendUserMessage(id, proto.CmdCreateAccount, &proto.CreateAccountSuccess{
			ErrCode: 0,
			Account: c.Account,
			Sign: "321321321",
		})
	} else {
		if sm.dbProxy.AddUserInfo(&proto.T_Users{
			Account: "guest_" + c.Account,
			Name: c.Name,
			Level: 1,
			Gems: 0,
			Exp: 0,
			Coins: 0,
			Sex: uint8(rand.Int()%2),
		}) {
			fmt.Println("create account add")
			sm.SendUserMessage(id, proto.CmdCreateAccount, &proto.CreateAccountSuccess{
				ErrCode: 0,
				Account: c.Account,
				Sign: "321321321",
			})
		} else {
			fmt.Println("create account add fail")
			sm.SendUserMessage(id, proto.CmdCreateAccount, &proto.CreateAccountSuccess{
				ErrCode: 1,
			})
		}
	}
}


