package main

import (
    "fmt"
    "net/http"
    "github.com/xingcuntian/go_test/go-example/redis/recordstore/models"
    "strconv"
)
//curl -i -L -d "id=2" localhost:3000/likes
func main(){
    http.HandleFunc("/album",showAlbum)
    http.HandleFunc("/likes",addLike)
    http.HandleFunc("/popular",listPopular)
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

func addLike(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST" {
        w.Header().Set("Allow","POST")
        http.Error(w,http.StatusText(405),405)
        return
    }

    id := r.PostFormValue("id")
    if id == "" {
        http.Error(w, http.StatusText(400),400)
        return
    }

    if _, err := strconv.Atoi(id); err != nil{
        http.Error(w,http.StatusText(400),400)
        return
    }

    err := models.IncrementLikes(id)
    if err == models.ErrNoAlbum {
        http.NotFound(w,r)
        return
    }else if err != nil{
        http.Error(w,http.StatusText(500),500)
        return
    }
    http.Redirect(w,r,"/album?id="+id,303)
}

func listPopular(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET" {
        w.Header().Set("Allow","GET")
        http.Error(w,http.StatusText(405),405)
        return
    }

    abs, err := models.FindTopThree()
    if err != nil {
        http.Error(w, http.StatusText(500),500)
        return
    }
    fmt.Println(abs)
    for i, ab := range abs {
        fmt.Println(ab)
        fmt.Fprintf(w, "%d) %s by %s: £%.2f [%d likes] \n", i+1, ab.Title, ab.Artist, ab.Price, ab.Likes)
    }
}
