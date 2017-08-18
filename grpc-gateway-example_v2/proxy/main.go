package main

import (
	"flag"
	"io"
	"net/http"

	"strings"

	"mime"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/go-bindata-assetfs"
	"github.com/xingcuntian/go_test/grpc-gateway-example_v2/pkg/ui/data/swagger"
	gw "github.com/xingcuntian/go_test/grpc-gateway-example_v2/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	echoEndpoint = flag.String("echo_endpoint", ":50052", "endpoint of yourService")
)

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	swmux := http.NewServeMux()
	swmux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(gw.Swagger))
	})

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}
	swmux.Handle("/", mux)
	serveSwagger(swmux)
	http.ListenAndServe(":7070", swmux)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
