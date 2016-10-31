package models

import (
   // "time"
    "github.com/xingcuntian/go_test/goapi/modules/base"
	"crypto/sha256"
    "errors"
	"fmt"
)

type User struct{
    Id  int64 `json:"id"`
    Username string `json:"username",xorm:"UNIQUE NOT NLL"`
    Email string `json:"email",xorm:"UNIQUE NOT NULL"`
    Phone string `json:"phone"`
    Password string `json:"_",xorm:"NOT NULL"`
   // RememberToken string `json:"remember_token"`
    //Salt string `json:"_",xorm:"NOT NULL"`
   // Created time.Time `json:"created",xorm:"CREATED"`
    //Updated time.Time `json:"updated",xorm:"UPDATED"`   
}
//创建用户
func CreateUser(uname,email,password,phone string)(int64,error){
    //u.Salt=GetUserSalt()
    //u.Password = EncodePassword(u.Password,u.Salt)
    user:=new(User)
    user.Username=uname
    user.Email=email
    user.Password=password
    user.Phone=phone
    _,err:=x.Insert(user)
    if err!=nil{
        return 0,err
    }
    return user.Id,nil
}
//生成salt
func GetUserSalt()string{
    return base.GetRandomString(10)
}
//用salt加密password
func EncodePassword(password,salt string)string{
    newPasswd:=base.PBKDF2([]byte(password),[]byte(salt),10000,50,sha256.New)
    return fmt.Sprintf("%x",newPasswd)
}

func GetUserInfo(uid int64)(user *User,err error){
    user=&User{Id:uid}
    has,err:=x.Get(user)
    if err!=nil{
        return
    }
    if !has{
        err=errors.New("user no found")
        return
    }
    return
}
