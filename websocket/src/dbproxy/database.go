package dbproxy

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

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

	return dp
}

func (dp *DbProxy) CreatTable(v ...interface{}) {
	dp.db.CreateTable(v...)
}

func (dp *DbProxy) CreateTableIfNot(v ...interface{}) {
	for _, m := range v {
		if dp.db.HasTable(m) == false {
			dp.db.CreateTable(m)
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



