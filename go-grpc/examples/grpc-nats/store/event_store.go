package store

import (
	pb "github.com/xingcuntian/go_test/go-grpc/examples/grpc-nats/order"
)

type EventStore struct{}

func (store EventStore) CreateEvent(order *pb.EventStore) error {
	session := mgoSession.Copy()
	defer session.Close()
	col := session.DB("xingcuntian").C("events")
	err := col.Insert(order)
	return err
}

func (store EventStore) GetEvents() []pb.EventStore {
	var events []pb.EventStore
	session := mgoSession.Copy()
	defer session.Close()
	col := session.DB("xingcuntian").C("events")
	iter := col.Find(nil).Iter()
	result := pb.EventStore{}
	for iter.Next(&result) {
		events = append(events, result)
	}
	return events
}
