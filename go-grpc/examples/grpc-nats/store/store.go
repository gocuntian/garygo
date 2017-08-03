package store

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

func init() {
	var err error
	mgoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"127.0.0.1"},
		Database: "xingcuntian",
		Username: "xingcuntian",
		Password: "123456",
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}
