package main

import (
    "log"
    _"github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

type User struct{
    Id int64
    Name string
}

var engine *xorm.Engine

func init(){
    var err error
    engine,err=xorm.NewEngine("mysql","xingcuntian:xingcuntian@2016@tcp(192.168.8.70:3306)/gotest?charset=utf8")
    if err!=nil{
        log.Fatalf("fail to create engine:%v\n",err)
    }
    err=engine.Sync2(new(User))
    if err!=nil{
        log.Fatalf("fail to sync data:$v\n",err)
    }
}

func CreateAccount(name string)error{
    _,err:=engine.Insert(&User{Name:name})
    return err
}

func UpdateAccount(name string,id int64)error{
    user:=new(User)
    user.Name=name
    _,err:=engine.Id(id).Update(user)
    return err  
}

func DeleteAccount(id int64)error{
    user:=new(User)
    _,err:=engine.Id(id).Delete(user)
    return err
}

func ListAccount()(list []User,err error){
    err=engine.Desc("id").Find(&list)
    return list,err
}