package main

import (
	"github.com/hashicorp/consul/api"
	"fmt"
)

func main(){
	//client, err := api.NewClient(api.DefaultConfig())
	client, err := api.NewClient(&api.Config{Address:"127.0.0.1:8500",Token:"3ea6e273-94c2-4e96-a518-c8fcd1bc4ae2"}) 
	if err != nil {
		panic(err)
	}

	kv := client.KV()
	//fmt.Println(*kv)
	// p := &api.KVPair{Key:"xingcuntian",Value:[]byte("test")}
	// _, err = kv.Put(p, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// pair, _, err := kv.Get("xingcuntian",nil)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pair)
	// fmt.Println(string(pair.Value))
	pair, _, _ := kv.Get("forthservice/key3",nil)
	fmt.Println(pair)
	fmt.Println("Key: ",pair.Key)
	fmt.Println("LockIndex: ",pair.LockIndex)
	fmt.Println("ModifyIndex: ",pair.ModifyIndex)
	fmt.Println("Session:", pair.Session)
	fmt.Println("CreateIndex:",pair.CreateIndex)
	fmt.Println("Flags:",pair.Flags)
	fmt.Println("value:",string(pair.Value))

	acl := client.ACL()

	// id ,_, err := acl.Create(&api.ACLEntry{
	// 	Name: "mtoken",
	// 	Type: "client",
	// 	Rules: "service \"firstservice\" {\n policy = \"write\"\n }\n key \"firstservice/\" {\n policy = \"read\"\n }",
	// },nil) //,&api.WriteOptions{Token:"master"}
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("create_token_id:",id)
	fmt.Println("=====================List============================================")
	list, _ , _ := acl.List(nil)
	fmt.Println(list)
	for k, val := range list {
		fmt.Println("key:",k,"obj: ",val)
	}
	fmt.Println()
	
	fmt.Println("====================Replication================================-===")
	status, _, err := acl.Replication(nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Println("Replication: ",status)
	fmt.Println("Replication.Enabled: ",status.Enabled)
	fmt.Println("Replication.ReplicatedIndex: ",status.ReplicatedIndex)
	fmt.Println("Replication.Running: ",status.Running)
	fmt.Println("Replication.SourceDatacenter: ",status.SourceDatacenter)
	fmt.Println("Replication.LastError: ",status.LastError)
	fmt.Println("Replication.LastSuccess: ",status.LastSuccess)
	fmt.Println("===================Info===================================")
	info, _, _ := acl.Info("e7b26393-54ac-2681-f248-f55c4a738e7c",nil)
	fmt.Println(info)
	// fmt.Println("====================Destroy is used to destroy a given ACL token ID=====")
	// _,err = acl.Destroy("aac48a4c-a0ad-d231-4a82-55edb6d14467",nil)
	// if err != nil {
	// 	fmt.Println("error: ", err.Error())
	// }
	// fmt.Println("=======destroy=============================")

	// fmt.Println("=============================Clone========================")
	// newId,WriteMeta,err := acl.Clone("7baa2200-2578-f8be-0403-668bca7d4a12",nil)
	// if err != nil {
	// 	fmt.Println("error: ", err.Error())
	// }
	// fmt.Println("NewId: ", newId)
	// fmt.Println("WriteMeta: ",WriteMeta)

	fmt.Println("=================Catalog ======================================")
	catalog := client.Catalog()
	cg, err := catalog.Datacenters()
	if err != nil {
		fmt.Println("error : ", err.Error())
	}
	fmt.Println(cg)
	// fmt.Println("=================Catalog.Deregister===========================")
	// _, err = catalog.Deregister(&api.CatalogDeregistration{
	// 	Node:"a4",
	// 	Address:"172.20.12.174",
	// 	Datacenter:"xingcuntian2",
	// 	//ServiceID:"gary1",
	// 	//CheckID:"gary002",
	// },nil)
	// if err != nil {
	// 	fmt.Println("error: ", err.Error())
	// }
	fmt.Println("=================Catalog.Node===========================")
	nodeInfo, _ , _ := catalog.Node("a1",nil)
	fmt.Println("node: ", nodeInfo)
	fmt.Println("node.Node: ",nodeInfo.Node)
	fmt.Println("node.services: ", nodeInfo.Services)

		fmt.Println("=================Catalog.Register===========================")
	_, err = catalog.Register(&api.CatalogRegistration{
		Node:"a4",
		Address:"172.20.12.174",
	},nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	
	
}

// &{32948 32948 e7b26393-54ac-2681-f248-f55c4a738e7c mytoken client service "forthservice" {
//  policy = "write"
//  }
//  key "forthservice/" {
//  policy = "read"
//  }}


// type ACLEntry struct {
//     CreateIndex uint64
//     ModifyIndex uint64
//     ID          string
//     Name        string
//     Type        string
//     Rules       string
// }

