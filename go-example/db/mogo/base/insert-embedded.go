package main

import (
	"fmt"
	"log"
	"time"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

type Task struct {
	Description string
	Due			time.Time
}

type Category struct {
	Id 			bson.ObjectId `bson:"_id,omitempty"`
	Name		string
	Description string
	Tasks		[]Task
}

func main(){
	session,err := mgo.Dial("mongodb://xingcuntian:123456@localhost:27017/xingcuntian")
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)
	c := session.DB("xingcuntian").C("Category")
	data1 := Category{
		bson.NewObjectId(),
		"Open Source",
		"Tasks for open-source projects",
		[]Task{
			Task{"Create project in mgo", time.Date(2015, time.August, 10, 0, 0, 0, 0, time.UTC)},
			Task{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
		},
	}
	err = c.Insert(&data1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Insert data1 ok")
	
	var count int
	count, err = c.Count()
	if err !=nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d records inserted", count)

	iter := c.Find(nil).Iter()
     result :=Category{}
	for iter.Next(&result) {
		fmt.Printf("Name:%s, Description:%s\n", result.Name, result.Description)
		tasks := result.Tasks
		for _, v := range tasks {
			fmt.Println(v.Description)
			fmt.Println(v.Due)
		}
	}

}