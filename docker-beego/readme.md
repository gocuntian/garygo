## 构建镜像

#### docker build -t go-web-demo .

## 创建并启动容器

#### docker run -it --name my-go-web-demo -p 8080:8080  -v go-web-demo:/go/src/go-web-demo -w /go/src/go-web-demo go-web-demo