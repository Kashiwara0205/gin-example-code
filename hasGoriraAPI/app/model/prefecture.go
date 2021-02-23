package model

import (
    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/mysql"
)

type Prefecture struct {
    gorm.Model
    Name string
}