package models

import (
    "fmt"
    _"github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/xingcuntian/go_test/goapi/modules/setting"
)

var (
    x *xorm.Engine
    tables []interface{}
    HasEngine bool
    DbCfg struct{
        Type,Host,Name,User,Pwd,Port string
    } 
)
func init(){
    tables = append(tables,new(User),new(Contact))
}

func LoadModelsConfig(){
    DbCfg.Type = setting.Cfg.MustValue("database","DB_TYPE")
    DbCfg.Host = setting.Cfg.MustValue("database","HOST")
    DbCfg.Port = setting.Cfg.MustValue("database","PORT")
    DbCfg.User = setting.Cfg.MustValue("database","USER")
    DbCfg.Pwd = setting.Cfg.MustValue("database","PASSWORD")
    DbCfg.Name = setting.Cfg.MustValue("database","NAME")
}

func getEngine()(*xorm.Engine,error){
    cnnstr:=""
    switch DbCfg.Type {
    case "mysql":
        cnnstr=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
        DbCfg.User,DbCfg.Pwd,DbCfg.Host,DbCfg.Port,DbCfg.Name)
    default:
    return nil,fmt.Errorf("Unknown database type:%s",DbCfg.Type)
   }
   return xorm.NewEngine(DbCfg.Type,cnnstr)
}

func NewEngine()(err error){
    //x,err:=getEngine() //this is error
    x,err:=getEngine()
    if err!=nil{
        return fmt.Errorf("models.init(fail to connect to database):%v",err)
    }
    if err = x.Sync2(tables...);err!=nil{
        return fmt.Errorf("Sync database struct error :%v\n",err)
    }
   return nil
}
