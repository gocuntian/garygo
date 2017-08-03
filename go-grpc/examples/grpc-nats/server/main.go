package main

import (
	//"context"
	"encoding/json"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	"github.com/satori/go.uuid"
	pb "github.com/xingcuntian/go_test/go-grpc/examples/grpc-nats/order"
	"github.com/xingcuntian/go_test/go-grpc/examples/grpc-nats/store"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port      = ":50051"
	aggregate = "Order"
	event     = "OrderCreated"
)

type server struct{}

func (s *server) CreateOrder(ctx context.Context, in *pb.Order) (*pb.OrderResponse, error) {
	store := store.OrderStore{}
	store.CreateOrder(in)
	go publishOrderCreated(in)
	return &pb.OrderResponse{IsSuccess: true}, nil
}

func publishOrderCreated(order *pb.Order) {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()
	eventData, _ := json.Marshal(order)
	event := pb.EventStore{
		AggregateId:   order.OrderId,
		AggregateType: aggregate,
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		EventData:     string(eventData),
	}
	subject := "Order.OrderCreated"
	data, _ := proto.Marshal(&event)
	natsConnection.Publish(subject, data)
	log.Println("Published message on subject " + subject)
}

func (s *server) GetOrders(filter *pb.OrderFilter, stream pb.OrderService_GetOrdersServer) error {
	store := store.OrderStore{}
	orders := store.GetOrders()
	for _, order := range orders {
		if err := stream.Send(&order); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	server := &server{}
	pb.RegisterOrderServiceServer(s, server)
	s.Serve(lis)
}
