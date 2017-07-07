package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var model BookmarkModel
var id string

func init() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"127.0.0.1"},
		Database: "xingcuntian",
		Username: "xingcuntian",
		Password: "123456",
		Timeout:  60 * time.Second,
	})
	// OR
	//session, err := mgo.Dial("mongodb://xingcuntian:123456@localhost:27017/xingcuntian")
	if err != nil {
		log.Fatalf("[MongoDB Session]: %s\n", err)
	}
	collection := session.DB("xingcuntian").C("bookmarks")
	//collection.RemoveAll(nil)
	model = BookmarkModel{
		C: collection,
	}
}

func create() {
	bookmark := Bookmark{
		Name:        "gorethink3",
		Description: "Go driver for RethinkDB",
		Location:    "https://github.com/dancannon/gorethink",
		Priority:    4,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "nosql", "rethinkdb"},
	}
	if err := model.Create(&bookmark); err != nil {
		log.Fatalf("[Create]: %s\n", err)
	}
	id := bookmark.ID.Hex()
	fmt.Printf("New bookmark has been inserted with ID: %s\n", id)
}

var wg sync.WaitGroup

func main() {
	// wg.Add(10000000)
	// fmt.Println("start")
	// go func() {
	// 	for i := 0; i < 10000000; i++ {
	// 		create()
	// 		wg.Done()
	// 	}
	// }()
	// wg.Wait()
	// fmt.Println("end")
	bookmarks := model.GetAll()
	fmt.Println(bookmarks)

	tags := []string{"mysql"}
	bookmarks = model.GetByTag(tags)
	fmt.Println(bookmarks)

	id := "595cbe0783cb616e670a74eb"
	bookmark, _ := model.GetByOID(id)
	fmt.Println(bookmark)
}
