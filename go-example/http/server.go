package main

import (
    "fmt"
    "net/http"
    "net/http/cgi"
)

func main(){
    http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
        handler:=new(cgi.Handler);
        handler.Path = "/usr/local/go/bin/go";
        script:="/data/golang/src/github.com/xingcuntian/go_test/go-example/http"+r.URL.Path;
        fmt.Println(handler.Path);
        handler.Dir = "/data/golang/src/github.com/xingcuntian/go_test/go-example/http";
        args:=[]string{"run",script};
        handler.Args = append(handler.Args,args...);
        fmt.Println(handler.Args);
        handler.ServeHTTP(w,r);
    });
    http.ListenAndServe(":8989",nil);
    select{}
}
