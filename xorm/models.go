package main

import (
    "errors"
    "log"
    _"github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

type User struct{
    Id int64
    Name string
}
type Contact struct{
    Name string `form:"name" binding:"Required"`
    Email string `form:"email"`
    Message string `form:"message" binding:"Required"`
    MailingAddress string `form:"mailing_address"`
}

var engine *xorm.Engine

func init(){
    var err error
    engine,err=xorm.NewEngine("mysql","xingcuntian:xingcuntian@2016@tcp(192.168.8.70:3306)/gotest?charset=utf8")
    if err!=nil{
        log.Fatalf("fail to create engine:%v\n",err)
    }
    err=engine.Sync2(new(User),new(Contact))
    if err!=nil{
        log.Fatalf("fail to sync data:%v\n",err)
    }
}

func CreateContact(contact Contact)error{
    _,err:=engine.Insert(contact)
    return err
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

func InfoAccount(id int64)(*User,error){
    one:=&User{}
    has,err:=engine.Id(id).Get(one)
    if err!=nil{
        return nil,err
    }else if !has{
        return nil,errors.New("User info not exist")
    }
    return one,nil
}