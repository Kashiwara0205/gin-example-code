package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "db"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}