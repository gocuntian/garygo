package account

import (
    //"encoding/json"
    m "github.com/xingcuntian/go_test/goapi/models"
    "github.com/xingcuntian/go_test/goapi/modules/middleware"
)
//创建用户
func Create(ctx *middleware.Context){
    uname:=ctx.Query("username")
    email:=ctx.Query("email")
    password:=ctx.Query("password")
    phone:=ctx.Query("phone")
    if len(uname) == 0 || len(email)==0 || len(password) ==0{
        ctx.ErrorJSON(401,"params error")
        return
    }
    //user:=new(m.User)
    //user.Username = uname
   // user.Email = email
   // user.Password = password
   // user.Phone = phone
   // {Username:uname,Email:email,Password:password,Phone:phone}
    uid,err:=m.CreateUser(uname,email,password,phone)
    if err!=nil || uid==0{
        ctx.ErrorJSON(502,"create user failed")
        return
    }
    ctx.SuccessJSON(uid)
    return
}
//用户信息
func Info(ctx *middleware.Context){
    uid:=ctx.QueryInt64("uid")
    if uid==0{
        ctx.ErrorJSON(401,"params errors")
        return
    }
    user,err:=m.GetUserInfo(uid)
    if err!=nil||user == nil{
        ctx.ErrorJSON(404,"user no found")
        return
    }
    ctx.SuccessJSON(user)
} 