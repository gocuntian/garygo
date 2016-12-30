package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var db *DB

func init(){
    var err error
    db,err = gorm.Open("mysql","xingcuntian:xingcuntian@2016@/gotest?charset=utf8&parseTime=True&loc=Local")
    defer db.Close()
}