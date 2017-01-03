package models

import (
	"log"
	 _"github.com/go-sql-driver/mysql"
    "github.com/xingcuntian/go-xorm/xorm"
    "github.com/xingcuntian/go-xorm/core"
)

type Todo struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Completed     bool   `json:"completed"`
	Created       []byte `json:"created"`
	Updated       []byte `json:"updated"`
	Due           []byte `json:"due"`
	DateCompleted []byte `json:"date_completed"`
}

type Todos []Todo

// apiError define structure of API error
type apiError struct {
	Tag     string `json:"-"`
	Error   error  `json:"-"`
	Message string `json:"error"`
	Code    int    `json:"code"`
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
        tbMapper:=core.NewPrefixMapper(core.SnakeMapper{},"t_")
        engine.SetTableMapper(tbMapper)
        //设置表后缀
        // tbMapper:=core.NewSuffixMapper(core.SnakeMapper{},"_suffix")
        // engine.SetTableMapper(tbMapper)
        //同步表结构
        err=engine.Sync2(new(Todo))
        if err!=nil{
            log.Fatalf("fail to sync data:%v\n",err)
        }
}
