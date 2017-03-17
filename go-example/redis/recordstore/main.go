package main

import (
    "fmt"
    "net/http"
    "github.com/xingcuntian/go_test/go-example/redis/recordstore/models"
    "strconv"
)

func main(){
    http.HandleFunc("/album",showAlbum)
    http.ListenAndServe(":3000",nil)
}

func showAlbum(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET" {
        w.Header().Set("Allow","GET")
        http.Error(w, http.StatusText(405),405)
        return
    }

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w,http.StatusText(400),400)
        return
    }

    if _, err := strconv.Atoi(id); err != nil{
        http.Error(w, http.StatusText(400),400)
        return
    }

    bk,err := models.FindAlbum(id)
    if err == models.ErrNoAlbum {
        http.NotFound(w,r)
        return
    }else if err !=nil {
        http.Error(w,http.StatusText(500),500)
        return
    }

    fmt.Fprintf(w, "%s by %s: £%.2f [%d likes] \n", bk.Title, bk.Artist, bk.Price, bk.Likes)
}