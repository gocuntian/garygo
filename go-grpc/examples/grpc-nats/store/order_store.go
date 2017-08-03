package store

import (
	pb "github.com/xingcuntian/go_test/go-grpc/examples/grpc-nats/order"
)

type OrderStore struct{}

func  (store OrderStore) CreateOrder(order *pb.Order) error {
	session := mgoSession.Copy()
	defer session.Close()
	col := session.DB("xingcuntan").C("orders")
	err := col.Insert(order)
	return err
}

func (store OrderStore) GetOrders() []pb.Order {
	var orders []pb.Order
	session := mgoSession.Copy()
	defer session.Close()
	col := session.DB("xingcuntian").C("orders")
	iter := col.Find(nil).Iter()
	result := pb.Order{}
	for iter.Next(&result) {
		orders = append(orders,result)
	}
	return orders
}