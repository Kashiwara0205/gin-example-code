package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
    _"github.com/jinzhu/gorm/dialects/mysql"
)

const(
    DbEngine = "mysql"
    DBUser = "user"
    DBPass = "password"
    DBProtocol = "tcp(db:3306)"
    DBName = "db"
)

func Connection() *gorm.DB {
    connectTemplate := "%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local"
    connectConfig := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
    db, err := gorm.Open(DbEngine, connectConfig)

    if err != nil {
        panic(err)
    }

    db.LogMode(true)
    return db
}