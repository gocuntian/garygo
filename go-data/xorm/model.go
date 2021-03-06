package main

import (
   // "errors"
    "time"
    "log"
    _"github.com/go-sql-driver/mysql"
    "github.com/xingcuntian/go-xorm/xorm"
    "github.com/xingcuntian/go-xorm/core"
)
type Member struct{
    Id int64
    Name string `xorm:"varchar(25) notnull unique 'user_name'"`
    Salt string
    Age int 
    Passwd string `xorm:"varchar(200)"`
    Created time.Time `xorm:"created"`
    Updated time.Time `xorm:"updated"`
    Version int64 `xorm:"version"`
}

type Group struct{
    Id int64
    Name string
}

type Type struct{
    Id int64
    Name string
}

type User struct{
    Id int64
    Name string
    GroupId int64 `xorm:"index"`
    TypeId int64  `xorm:"index"`
}



// 1)
// type UserGroup struct{
//     User `xorm:"extends"`
//     Name string
// }

//[
//  {{1 test1 1} group1} 
//  {{2 test2 2} group2}
//]

type UserGroup struct{
    User `xorm:"extends"`
    Group `xorm:"extends"`
}
// [
//  {{1 test1 1} {1 group1}} 
//  {{2 test2 2} {2 group2}}
// ]



type UserGroupType struct{
    User `xorm:"extends"`
    Group `xorm:"extends"`
    Type `xorm:"extends"`
}
//[
// {{1 test1 1 1001} {1 group1} {1001 person}}
// {{2 test2 2 1002} {2 group2} {1002 super}}
//]


func (UserGroup) TableName()string{
    return "prefix_user"
}
var engine *xorm.Engine

func init(){
        var err error 
        //链接数据库
        engine,err=xorm.NewEngine("mysql","xingcuntian:xingcuntian@2016@tcp(192.168.8.70:3306)/gotest?charset=utf8")
        if err!=nil{
            log.Fatalf("fail to create engine:%v\n",err)
        }
        //设置标前缀
        tbMapper:=core.NewPrefixMapper(core.SnakeMapper{},"prefix_")
        engine.SetTableMapper(tbMapper)
        //设置表后缀
        // tbMapper:=core.NewSuffixMapper(core.SnakeMapper{},"_suffix")
        // engine.SetTableMapper(tbMapper)
        //同步表结构
        err=engine.Sync2(new(Member),new(User),new(Group),new(Type))
        if err!=nil{
            log.Fatalf("fail to sync data:%v\n",err)
        }
}

func InsertDb(member *Member)error{
   _,err:=engine.Insert(member)
   return err
}

func InsertBatch(member []Member) error{
    _,err:=engine.Insert(&member)
    return err
}

func InsertMore(member []*Member)error{
    _,err:=engine.Insert(&member)
    return err
}