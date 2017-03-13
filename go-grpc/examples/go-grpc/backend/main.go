package main

import (
    "flag"
    "fmt"
    "log"
    "math/rand"
    "net"
    "net/http"
    _ "net/http/pprof"
    "time"
    "golang.org/x/net/context"
    "golang.org/x/net/trace"
    pb "github.com/xingcuntian/go_test/go-grpc/examples/go-grpc/search"
    "google.golang.org/grpc"
)

var (
    index = flag.Int("index",0,"RPC port is 36061+index; debug port is 36661+index")
)

type server struct{}

func randomDuration(max time.Duration) time.Duration{
    return time.Duration(1+int64(rand.Intn(10))) * (max / 10)
}

func logSleep(ctx context.Context, d time.Duration){
    if tr, ok := trace.FromContext(ctx); ok{
        tr.LazyPrintf("sleeping for %s",d)
    }
}

func (s *server) Search(ctx context.Context, req *pb.Request) (*pb.Result, error) {
    d := randomDuration(100 * time.Millisecond)
    logSleep(ctx,d)
    select{
        case <-time.After(d):
            return &pb.Result{
                Id:1,
                Username:"test",
                Avatar:"http://image.bi.sensetime.com",
                CompanyId:2,
            },nil
        case <-ctx.Done():
            return nil,ctx.Err()    
    }
}


func (s *server) Watch(req *pb.Request, stream pb.Google_WatchServer) error {
    ctx := stream.Context()
   for i := 0; ; i++ {
        d := randomDuration(1 * time.Second)
        logSleep(ctx,d)
        select{
            case <-time.After(d):
                err := stream.Send(&pb.Result{
                    Id:1,
                    Username:"test",
                    Avatar:"http://image.bi.sensetime.com",
                    CompanyId:2,
                })
                if err != nil {
                    return err
                }
            case <-ctx.Done():
                    return ctx.Err()
        }
    }
}

func main(){
    flag.Parse()
    rand.Seed(time.Now().UnixNano())
    go http.ListenAndServe(fmt.Sprintf(":%d",36661+*index),nil)
    lis,err := net.Listen("tcp",fmt.Sprintf(":%d",36061+*index))
    if err != nil {
        log.Fatalf("failed to listen:%v",err)
    }
    g := grpc.NewServer()
    pb.RegisterGoogleServer(g,new(server))
    g.Serve(lis)
}
