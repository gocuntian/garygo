package rpcexample

import (
    "log"
)
//在RPC调用中保存要传递给服务的参数Arith
type Args struct{
    A, B int
}
//服务方法
type Arith int
//RPC调用的结果是这种类型的
type Result int

func (t *Arith) Multiply(args Args, result *Result) error {
    return Multiply(args,result)
}

func Multiply(args Args, result *Result) error{
    log.Printf("Multiply %d with %d\n",args.A,args.B)
    *result = Result(args.A * args.B)
    return nil
}
