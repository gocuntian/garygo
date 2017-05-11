package models

import (
    "fmt"
    "errors"
    "github.com/mediocregopher/radix.v2/pool"
    "github.com/mediocregopher/radix.v2/redis"
    "log"
    "strconv"
)

var db *pool.Pool

func init(){
    var err error
    db, err = pool.New("tcp","localhost:6379",10)
    if err != nil {
        log.Panic(err)
    }
}

var ErrNoAlbum = errors.New("models: no album found")

type Album struct{
    Title string
    Artist string
    Price float64
    Likes int
}

func populateAlbum(reply map[string]string)(*Album, error){
    var err error
    ab := new(Album)
    ab.Title = reply["title"]
    ab.Artist = reply["artist"]
    ab.Price, err = strconv.ParseFloat(reply["price"],64)
    if err != nil {
        return nil, err
    }
    ab.Likes, err = strconv.Atoi(reply["likes"])
    if err != nil {
        return nil, err
    }
    return ab, nil
}

func FindAlbum(id string)(*Album, error){
    conn, err := db.Get()
    if err != nil {
        return nil, err
    }
    defer db.Put(conn)

    err = conn.Cmd("auth","sensetime@2016").Err;
    if err != nil {
        return nil, err
    }

    reply, err := conn.Cmd("HGETALL","album:"+id).Map()
    fmt.Println(reply)
    if err != nil{
        return nil, err
    }else if len(reply) == 0 {
        return nil, ErrNoAlbum
    }
    return populateAlbum(reply)
}

func IncrementLikes(id string) error {
    conn, err := db.Get()
    if err != nil {
        return err
    }
    defer db.Put(conn)

    err = conn.Cmd("auth","sensetime@2016").Err
    if err != nil {
        return err
    }

    exists, err := conn.Cmd("EXISTS","album:"+id).Int()
    if err != nil {
        return err
    }else if exists == 0{
        return ErrNoAlbum
    }
   //事务块的开始 multi return ok
    err = conn.Cmd("MULTI").Err
    if err != nil{
        return err
    }
    err = conn.Cmd("HINCRBY","album:"+id,"likes",1).Err
    if err != nil {
        return err
    }

    err = conn.Cmd("ZINCRBY", "likes", 1, id).Err
    if err != nil {
        return err
    }
    //事务块内所有命令的返回值，按命令执行的先后顺序排列。 
    //当操作被打断时，返回空值 nil 。
    err = conn.Cmd("EXEC").Err
    if err != nil {
        return err
    }

    return nil
}

func FindTopThree()([]*Album,error) {
    conn, err := db.Get()
    if err != nil {
        return nil, err
    }
    defer db.Put(conn)

    err = conn.Cmd("auth","sensetime@2016").Err
    if err != nil {
        return nil, err
    }

    for {
        err = conn.Cmd("WATCH", "likes").Err
        if err != nil {
            return nil, err
        }
 
        reply, err := conn.Cmd("ZREVRANGE", "likes", 0, 2).List()
        fmt.Println(reply)
        if err != nil {
            return nil, err
        }
 
        err = conn.Cmd("MULTI").Err
        if err != nil {
            return nil, err
        }
 
        for _, id := range reply {
            err := conn.Cmd("HGETALL", "album:"+id).Err
            if err != nil {
                return nil, err
            }
        }
 
        ereply := conn.Cmd("EXEC")
        fmt.Println(ereply)
        if ereply.Err != nil {
            return nil, err
        } else if ereply.IsType(redis.Nil) {
            continue
        }
 
        areply, err := ereply.Array()
         fmt.Println(areply)
        if err != nil {
            return nil, err
        }
 
        abs := make([]*Album, 3)
 
        for i, reply := range areply {
            mreply, err := reply.Map()
            fmt.Println(mreply)
            if err != nil {
                return nil, err
            }
            ab, err := populateAlbum(mreply)
            if err != nil {
                return nil, err
            }
            abs[i] = ab
        }
 
        return abs, nil
    }
}

