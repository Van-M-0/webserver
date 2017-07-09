package proto

/*
import "reflect"

func regiseter(i interface{}) reflect.Type {
	return reflect.TypeOf(i).Elem()
}

func NewStruct(cmd uint16) (interface{}, bool) {
	if i, ok := messageCenter[cmd]; ok {
		return reflect.New(i).Elem().Interface(), true
	}
	return nil, false
}

var messageCenter = map[uint16]reflect.Type{
	CmdGuestAuth: 		regiseter((*AuthGuest_C)(nil)),
	CmdLoginLobby:		regiseter((*Login_C)(nil)),
	CmdCreateAccount: 	regiseter((*CreateAccount_C)(nil)),
}
*/

func NewStruct(cmd uint16) (interface{}, bool) {
	switch cmd {
	case CmdGuestAuth:
		return &AuthGuest_C{}, true
	case CmdLoginLobby:
		return &Login_C{}, true
	case CmdCreateAccount:
		return &CreateAccount_C{}, true
	case CmdCreateRoom:
		return &CreateRoomXZ_C{}, true
	}
	return nil, false
}

const (
	CmdWebgateStart 		uint16 = 101
	CmdLoginStart 			uint16 = 201
	CmdGameStart 			uint16 = 1001
	CmdInvalid 				uint16 = 2001
)

const (
	// common proto
	CmdInvalidMessage 			uint16 = 1

	// web gate way protocol 101 - 200
	CmdGetServerInfo 			uint16 = 101

	// login server proto 201 - 300
	CmdGuestAuth				uint16 = 201
	CmdLoginLobby 				uint16 = 202
	CmdCreateAccount 			uint16 = 203
	CmdCreateRoom 				uint16 = 204

	//game server proto 1001 - 2000

	//server-server proto
	CmdRegisterServer 			uint16 = 3001

)

// server - server  proto
type RegiseterServer struct {

}

type ServerInfo struct {
	Name 		string				`json:"name"`
	Type 		int 				`json:"type"`
	Id 			int					`json:"id"`
}

type ServerList struct {
	Servers 	[]*ServerInfo		`json:"servers"`
}

type Message struct {
	Cmd 		uint16				`json:"cmd"`
	Msg 		interface{}			`json:"msg"`
}

type AuthGuest_C struct {
	Account 	string 				`json:"account"`
}

type GuestAuthMessage struct {
	ErrCode 	int 				`json:"errcode"`
	Account 	string				`json:"account"`
	Sign 		string 				`json:"sign"`
}

type Login_C struct {
	Account 	string				`json:"account"`
	Sign 		string				`json:"sign"`
}

type LoginLobbyScucess struct {
	ErrCode 	int 				`json:"errcode"`
	Account 	string 				`json:"account"`
	UserId 		uint32 				`json:"userid"`
	Name 		string 				`json:"name"`
	Level 		uint8 				`json:"lv"`
	Exp 		uint32 				`json:"exp"`
	Coins 		uint32 				`json:"coins"`
	Gems 		uint32 				`json:"gems"`
	RoomId 		string 				`json:"roomid"`
	Sex 		uint8 				`json:"sex"`
	Ip 			string  			`json:"ip"`
}

type CreateAccount_C struct {
	Account 	string 				`json:"account"`
	Sign 		string 				`json:"sign"`
	Name 		string 				`json:"name"`
}

type CreateAccountSuccess struct {
	ErrCode 	int 				`json:"errcode"`
	Account 	string				`json:"account"`
	Sign 		string				`json:"sign"`
}

type XZRoomConf struct {
	CellScore 	int					`json:"difen"`
	Zimo 		int					`json:"zimo"`
	Jiangdui 	bool				`json:"jiangdui"`
	Huansanzhang bool				`json:"huansanzhang"`
	ZuidaFan 	int					`json:"zuidafanshu"`
	Jushu 		int					`json:"jushuxuanze"`
	Dianganghua int				`json:"dianganghua"`
	Menqing 	bool				`json:"menqing"`
	Tiandihu 	bool				`json:"tiandihu"`
	Type 		string				`json:"type"`
}

type CreateRoomXZ_C struct {
	Account 	string 				`json:"account"`
	Sign 		string				`json:"sign"`
	Conf		XZRoomConf			`json:"conf"`
}

type EnterRoom_C struct {
	Account 	string 				`json:"account"`
	Sign 		string 				`json:"sign"`
	RoomId 		string 				`json:"roomid"`
}
