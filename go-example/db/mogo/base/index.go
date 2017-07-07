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
	c := session.DB("xingcuntian").C("categories")
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		log.Fatalln(err)
	}
	err = c.Insert(
		&Category{bson.NewObjectId(), "xingcuntian2", "this is username2"},
		&Category{bson.NewObjectId(), "linzihao2", "this is uname2"},
		&Category{bson.NewObjectId(), "huangxiang2", "this is students2"},
	)
	if err != nil {
		log.Fatalln(err)
	}
	result := Category{}
	err = c.Find(bson.M{"name": "xingcuntian2"}).One(&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Description:", result.Description)
	}
	var b []Category
	result2 := Category{}
	iter := c.Find(nil).Iter()
	for iter.Next(&result2) {
		b = append(b, result2)
	}
	fmt.Println(b)
}
