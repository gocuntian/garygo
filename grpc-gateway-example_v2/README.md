
# grpc_gateway + swagger


## https://coreos.com/blog
## https://coreos.com/blog/grpc-protobufs-swagger.html

# swagger UI
## swagger 集成
 #### pb 目录下生成一个 helloworld.swagger.json 文件。我们在 pb 目录下直接新增一个文件 helloworld.swagger.go，然后在里面定义一个常量 Swagger，内容即为 helloworld.swagger.json 的内容。
 #### 修改 proxy.go 文件中的 run() 方法来添加一个 API 路由来返回 swagger.json 的内容，
 ####  http://localhost:8080/swagger.json 便得到 helloworld RESTful API 的 swagger 说明了。 swagger.json 内容显示太不直观了。swagger 提供了非常好的可视化 swagger-ui。我们将 swagger-ui 添加到我们的 gateway 中。

## 下载 swagger-ui 代码

 #### Swagger 提供了可视化的 API 说明。我们可以在 RESTful JSON API gateway 中添加 swagger-ui。将 Swagger 源码的 dist 目录下包含了 swagger ui 所需的 HTML、css 和 js 代码文件，我们将该目录下的所有文件拷贝到 third_party/swagger-ui 目录下。将 swagger-ui 文件制作成 go 内置文件 我们可以使用 go-bindata 将 swagger-ui 的文件制作成 go 内置的数据文件进行访问。

## 先安装 go-bindata，

 #### $ go get -u github.com/jteeuwen/go-bindata/...
## 然后将 third-party/swagger-ui 下的所有文件制作成 go 内置数据文件，

 #### $ go-bindata --nocompress -pkg swagger -o pkg/ui/data/swagger/datafile.go third_party/swagger-ui/...
##生成文件 pkg/ui/data/swagger/datafile.go，

## $ls -l pkg/ui/data/swagger/datafile.go 
## swagger-ui 文件服务器

 #### 使用 go-bindata 将 swagger-ui 制作成 go 内置数据文件之后，我们便可以使用 elazarl/go-bindata-assetfs 结合 net/http 来将 swagger-ui 内置数据文件对外提供服务。

##安装 elazarl/go-bindata-assetfs，

## $ go get github.com/elazarl/go-bindata-assetfs/...
#### 默认打开的是一个 http://petstore.swagger.io/v2/swagger.json 的 API 说明信息。我们需要在输入框中输入我们 API 的地址 http://localhost:8080/swagger.json ，然后点击回车键才能看到我们的 API 说明，
#### 将文件 third_party/swagger-ui/index.html 中的 http://petstore.swagger.io/v2/swagger.json 替换成 http://localhost:8080/swagger.json ，然后重新生成 pkg/ui/data/swagger/datafile.go 文件，再重新编译一下即可。

## go get github.com/philips/go-bindata-assetfs

