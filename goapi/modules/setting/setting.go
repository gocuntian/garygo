package setting

import (
    "github.com/xingcuntian/goconfig"
    "os"
    "log"
    "os/exec"
    "path"
    "path/filepath"
    "strings"
)
var (
    Cfg *goconfig.ConfigFile
    Domain string
    HttpAddr,HttpPort string
    //CustomPath string
    ConfRootPath string
)

func ExecPath()(string,error){
    file, err := exec.LookPath(os.Args[0])
    if err!=nil{
        return "",err
    }
    p,err:=filepath.Abs(file)
    if err!=nil{
        return "",err
    }
    return p,nil
}

func WorkDir()(string,error){
    execPath, err := ExecPath()
    return path.Dir(strings.Replace(execPath, "\\", "/", -1)), err

}

func NewConfigContext(){
    workDir,err:=WorkDir()
    if err!=nil{
        log.Fatal(4,"Fail to get work directory:%v",err)
    }
    ConfRootPath = path.Join(strings.Replace(workDir,"/bin","/",-1),"src/github.com/xingcuntian/go_test/goapi/conf")
    Cfg,err=goconfig.LoadConfigFile(path.Join(ConfRootPath,"/app.ini"))
    if err!=nil{
        log.Fatal(4,"Fail to parse 'conf/app.ini':%v",ConfRootPath)
    }
    Domain = Cfg.MustValue("server","DOMAIN","localhost")
    HttpAddr = Cfg.MustValue("server","HTTP_ADDR","0.0.0.0")
    HttpPort = Cfg.MustValue("server","HTTP_PORT","3000")
}
