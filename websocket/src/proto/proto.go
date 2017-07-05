package proto

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

	//game server proto 1001 - 2000

	//server-server proto
	CmdRegisterServer 			uint16 = 3001

)

// server - server  proto
type RegiseterServer struct {

}

type ServerInfo struct {
	Name 		string		`codec:"name"`
	Type 		int 		`codec:"type"`
	Id 			int			`codec:"id"`
}

type ServerList struct {
	Servers 	[]*ServerInfo	`codec:"servers"`
}

type Message struct {
	Cmd 		uint16				`codec:"cmd"`
	Msg 		interface{}			`codec:"msg"`
}

type GuestAuthMessage struct {
	Account 	string				`codec:"account"`
}



