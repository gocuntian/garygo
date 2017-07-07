package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
}

func main() {
	session, err := mgo.Dial("mongodb://xingcuntian:123456@localhost:27017/xingcuntian")
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("xingcuntian").C("category2")
	data := Category{
		bson.NewObjectId(),
		"xingcuntian33",
		"this is test3333",
	}
	err = c.Insert(&data)
	if err != nil {
		log.Fatalln(err)
	}
	//insert two category objects
	err = c.Insert(&Category{bson.NewObjectId(), "xingcuntian444", "this is test4444"}, &Category{bson.NewObjectId(), "xingcuntian5555", "this is test5555"})
	if err != nil {
		log.Fatalln(err)
	}
	var count int
	count, err = c.Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count:", count)
}
