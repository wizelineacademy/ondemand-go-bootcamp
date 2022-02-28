package datastore

import (
	"log"

	// _ "github.com/PasHdez/ondemand-go-bootcamp/config"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDB() *gorm.DB {
	// DBMS := "mysql"
	// mySqlConfig := &mysql.Config{
	// 	User:                 config.C.Database.User,
	// 	Passwd:               config.C.Database.Password,
	// 	Net:                  config.C.Database.Net,
	// 	Addr:                 config.C.Database.Addr,
	// 	DBName:               config.C.Database.DBName,
	// 	AllowNativePasswords: config.C.Database.AllowNativePasswords,
	// 	Params: map[string]string{
	// 		"parseTime": config.C.Database.Params.ParseTime,
	// 	},
	// }

	db, err := gorm.Open("sqlite3", "gorm.db")

	//db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
