package routers

import (
    "github.com/xingcuntian/go_test/goapi/models"
    "github.com/xingcuntian/go_test/goapi/modules/setting"
	"log"
)

func GlobalInit(){
    setting.NewConfigContext()
    models.LoadModelsConfig()
    if err:=models.NewEngine();err!=nil{
        log.Fatal("fail to initialize orm engine:%v",err)
    }
    models.HasEngine=true
}