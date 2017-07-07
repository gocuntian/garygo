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
		log.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("xingcuntian").C("category2")
	data := map[string]string{"name": "xingcuntian", "description": "this is test"}
	err = c.Insert(data)
	if err != nil {
		log.Fatalln(err)
	}

	data1 := bson.D{
		{"name", "xingcuntian2"},
		{"description", "this is test2"},
	}
	fmt.Println(data1)
	err = c.Insert(data1)
	if err != nil {
		log.Fatalln(err)
	}
	var count int
	count, err = c.Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Count:", count)
}
