package main

import (
    "sync"
    "net/http"
)

func main(){
    var wg sync.WaitGroup
    var urls = []string{
        "http://www.golang.org/",
        "http://www.google.com/",
        "http://www.baidu.com/",
    }
    for _, url := range urls {
        wg.Add(1)
        go func(url string){
            defer wg.Done()
            http.Get(url)
        }(url)
    }
    wg.Wait()
}