package middleware

import (
    "gopkg.in/macaron.v1"
)

type Context struct{
    *macaron.Context
}

type resMsg struct{
    Code int `json:"code"`
    Data interface{} `json:"data"`
}
//返回错误信息
func (ctx *Context) ErrorJSON(code_no int,msg string){
    ctx.Context.Render.JSON(200,&resMsg{Code:code_no,Data:msg})
}
//返回正确信息
func (ctx *Context) SuccessJSON(data interface{}){
    ctx.Context.Render.JSON(200,&resMsg{Code:200,Data:data})
}
func Contexter() macaron.Handler{
    return func(c *macaron.Context){
        ctx:=&Context{
            Context:c,
        }
        c.Map(ctx)
    }
}