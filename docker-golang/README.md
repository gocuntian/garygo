go build -o main .

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

docker build -t xingcuntian/go-scratch -f Dockerfile.scratch .

docker run --name goscatch -d -p 8080:8080 xingcuntian/go-scratch
