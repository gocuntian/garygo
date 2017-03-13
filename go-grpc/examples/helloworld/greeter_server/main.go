package main

import (
	"log"
	"net"
	"fmt"
	"strings"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/xingcuntian/go_test/go-grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
	"github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

const (
	port = ":50051"
)


type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
     db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bi_admin_sg?charset=utf8&parseTime=True&loc=Local")
	 if err != nil{
        log.Fatalf("connect db failed: %v", err)
	 }
	 defer db.Close()
	 var result pb.HelloReply
     //var users []pb.HelloReply
	// db.Where([]int64{20, 21, 22}).Find(&users)
    //  fmt.Println(users)
    //return &users,nil
	 db.Raw("SELECT id, username,avatar,company_id FROM t_user WHERE id = ?", in.Id).Scan(&result)
	//fmt.Println(result)
	 return &result,nil
}




func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Greeter_GetCustomersServer) error{
	 db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bi_admin_sg?charset=utf8&parseTime=True&loc=Local")
	 if err != nil{
        log.Fatalf("connect db failed: %v", err)
	 }
	 defer db.Close()
	 var customers []*pb.HelloReply
	// db.Where([]int64{20, 21, 22}).Find(&customers)
	 db.Find(&customers)
	 fmt.Println(customers)
	 for _,customer := range customers {
		 if filter.Keyword != "" {
			 if !strings.Contains(customer.Username, filter.Keyword) {
				continue
			}
		 }
		 if err := stream.Send(customer); err != nil {
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
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
