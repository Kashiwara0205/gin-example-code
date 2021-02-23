package main

import (
    "hasGoriraAPI/db"
    "hasGoriraAPI/app/model"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
    db := db.Connection()
    defer db.Close()

    db.AutoMigrate(&model.Prefecture{})
}