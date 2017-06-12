package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"127.0.0.1:2379"}
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	log.Println("存储值")
	if ret, err := cli.Put(context.TODO(), "userinfo", `{user01:{topic:"xingcuntian"}}`); err != nil {
		log.Fatal(err)
	} else {
		log.Println("ret: ", ret)
	}

	log.Println("获取值")
	if resp, err := cli.Get(context.TODO(), "userinfo"); err != nil {
		log.Fatal(err)
	} else {
		log.Println("resp: ", resp)
	}

	log.Println("事务&超时")
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Txn(ctx).
		If(clientv3.Compare(clientv3.Value("key"), ">", "abc")). // txn value comparisons are lexical
		Then(clientv3.OpPut("key", "XYZ")).                      // this runs, since 'xyz' > 'abc'
		Else(clientv3.OpPut("key", "ABC")).
		Commit()
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("监视")

	rch := cli.Watch(context.Background(), "", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

}
