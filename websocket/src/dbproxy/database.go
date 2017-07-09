package dbproxy

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"proto"
)

//CREATE DATABASE IF NOT EXISTS mygame default charset utf8 COLLATE utf8_general_ci;
type DbProxy struct {
	Opts 		*DbOption
	uri 		string
	db 			*gorm.DB
}

type DbOption struct {
	Host		string
	Name 		string
	User 		string
	Pass 		string
	TablePrefix string
	ShowDetailLog bool
	Singular	bool
}

func NewDbProxy(opt *DbOption) *DbProxy {
	dp := &DbProxy{
		Opts: opt,
	}
	uri := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True",
		opt.User,
		opt.Pass,
		opt.Host,
		opt.Name,
	)
	fmt.Println("db proxy connection info ", uri)
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		fmt.Println("create db proxy err ", err)
		return nil
	}

	if opt.ShowDetailLog {
		db.LogMode(true)
	}

	if opt.Singular {
		db.SingularTable(true)
	}

	dp.db = db

	dp.InitTable()

	return dp
}

func (dp *DbProxy) CreateTable(v ...interface{}) {
	dp.db.CreateTable(v...)
}

func (dp *DbProxy) CreateTableIfNot(v ...interface{}) {
	for _, m := range v {
		if dp.db.HasTable(m) == false {
			dp.db.CreateTable(m).Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
		}
	}
}

func (dp *DbProxy) CreateTableForce(v...interface{}) {
	dp.db.DropTableIfExists(v...)
	dp.db.CreateTable(v...)
}

func (dp *DbProxy) DropTable(v ...interface{}) {
	dp.db.DropTableIfExists(v...)
}

// logic handler

func (dp *DbProxy) InitTable() {
	dp.CreateTableIfNot(&proto.T_Accounts{})
	dp.CreateTableIfNot(&proto.T_Games{})
	dp.CreateTableIfNot(&proto.T_GamesArchive{})
	dp.CreateTableIfNot(&proto.T_Guests{})
	dp.CreateTableIfNot(&proto.T_Message{})
	dp.CreateTableIfNot(&proto.T_Rooms{})
	dp.CreateTableIfNot(&proto.T_RoomUser{})
	dp.CreateTableIfNot(&proto.T_Users{})
	dp.CreateTableIfNot(&proto.T_MyTest{})
}

func (dp *DbProxy) PreLoadData() {

}

// t_accounts : account info
func (dp *DbProxy) GetAccountInfo(account string, accInfo *proto.T_Accounts) bool {
	return dp.db.Where(&proto.T_Accounts{Account: account}).First(accInfo).RowsAffected != 0
}

func (dp *DbProxy) AddAccountInfo(accInfo *proto.T_Accounts) bool {
	return dp.db.Create(accInfo).RowsAffected != 0
}

// t_users : user info
func (dp *DbProxy) AddUserInfo(userInfo *proto.T_Users) bool {
	fmt.Println("add user info : ", userInfo)
	return dp.db.Create(userInfo).RowsAffected != 0
}

func (dp *DbProxy) GetUserInfo(account string, userInfo *proto.T_Users) bool {
	return dp.db.Where("account = ? ", account).
		Select("userid, account, name, sex, headimg, level, exp, coins, gems, roomid").
		Find(&userInfo).
		RowsAffected != 0
}

func (dp *DbProxy) GetUserInfoByUserid(userid uint32, userInfo *proto.T_Users) bool {
	return dp.db.Where("userid = ? ", userid).
		Select("userid, account, name, sex, headimg, level, exp, coins, gems, roomid").
		Find(&userInfo).
		RowsAffected != 0
}

func (dp *DbProxy) ModifyUserInfo(userid uint32, userInfo *proto.T_Users) bool {
	return dp.db.Model(&proto.T_Users{}).
		Where("userid = ?", userid).
		Update(userInfo).
		RowsAffected != 0
}

func (dp *DbProxy) GetUserHistoryByUserid(userid uint32, userInfo *proto.T_Users) bool {
	return dp.db.Where("userid = ? ", userid).
		Select("history").
		Find(&userInfo).
		RowsAffected != 0
}

func (dp *DbProxy) GetUserGemsByUserid(userid uint32, userInfo *proto.T_Users) bool {
	return dp.db.Where("userid = ? ", userid).
		Select("gems").
		Find(&userInfo).
		RowsAffected != 0
}

func (dp *DbProxy) GetUserBaseInfo(userid uint32, userInfo *proto.T_Users) bool {
	return dp.db.Where("userid = ? ", userid).
		Select("name, sex, headimg").
		Find(&userInfo).
		RowsAffected != 0
}

// t_rooms : room info
func (dp *DbProxy) GetRoomInfo(roomid string, roomInfo *proto.T_Rooms) bool {
	return dp.db.Where(&proto.T_Rooms{Id: roomid}).First(roomInfo).RowsAffected != 0
}
