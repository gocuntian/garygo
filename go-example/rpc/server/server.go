package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"

	"github.com/xingcuntian/go_test/go-example/rpc/contract"
)

const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {
	helloWorld := new(HelloWorldHandler)
	rpc.Register(helloWorld)
	userHandler := &UserHandler{}
	rpc.Register(userHandler)
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()
	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

type UserHandler struct{}

func (h *UserHandler) GetUser(args *contract.UserRequest, reply *contract.UserResponse) error {
	reply.Name = "xingcuntian by id" + strconv.Itoa(int(args.Id))
	return nil
}
