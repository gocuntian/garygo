package main

import (
    "fmt"
    //"net/url"
    "net/http"
)

type MyHandler struct{}


// type ResponseWriter interface {

// Header() Header
// Header返回一个Header类型值，该值会被WriteHeader方法发送。 
// 在调用WriteHeader或Write方法后再改变该对象是没有意义的。


// WriteHeader(int)
// WriteHeader该方法发送HTTP回复的头域和状态码。 
// 如果没有被显式调用，第一次调用Write时会触发隐式调用WriteHeader(http.StatusOK) 
// WriterHeader的显式调用主要用于发送错误码。 


      
  
  // Write([]byte) (int, error)
  // Write向连接中写入作为HTTP的一部分回复的数据。  
  // 如果被调用时还未调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK) 
  // 如果Header中没有"Content-Type"键， 
  // 本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。 
   
// }

/*type Request struct {
    // Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET。
    Method string
   // URL在服务端表示被请求的URI，在客户端表示要访问的URL。 // // 在服务端，URL字段是解析请求行的URI（保存在RequestURI字段）得到的，
     // 对大多数请求来说，除了Path和RawQuery之外的字段都是空字符串。
    // 在客户端，URL的Host字段指定了要连接的服务器， // 而Request的Host字段（可选地）指定要发送的HTTP请求的Host头的值。
    URL *url.URL

    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0

    //	Host: example.com
    //	accept-encoding: gzip, deflate
    //	Accept-Language: en-us
    //	fOO: Bar
    //	foo: two
    //
    // then
    //
    //	Header = map[string][]string{
    //		"Accept-Encoding": {"gzip, deflate"},
    //		"Accept-Language": {"en-us"},
    //		"Foo": {"Bar", "two"},
    //	}
    //
    Header Header

    Body io.ReadCloser

    GetBody func() (io.ReadCloser, error)

    ContentLength int64

    TransferEncoding []string

    Close bool

    Host string
//Form是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据。 // 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。 
    Form url.Values
// PostForm是解析好的POST或PUT的表单数据。 // 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。 
    PostForm url.Values

// MultipartForm是解析好的多部件表单，包括上传的文件。 // 本字段只有在调用ParseMultipartForm后才有效。 // 在客户端，会忽略请求中的本字段而使用Body替代。 
    MultipartForm *multipart.Form

// Trailer指定了会在请求主体之后发送的额外的头域。 
// 在服务端，Trailer字段必须初始化为只有trailer键，所有键都对应nil值。 
// （客户端会声明哪些trailer会发送） \
// 在处理器从Body读取时，不能使用本字段。
// 在从Body的读取返回EOF后，Trailer字段会被更新完毕并包含非nil的值。 
// （如果客户端发送了这些键值对），此时才可以访问本字段。 
// // 在客户端，Trail必须初始化为一个包含将要发送的键值对的映射。（值可以是nil或其终值） 
// ContentLength字段必须是0或-1，以启用"chunked"传输编码发送请求。 /
/ 在开始发送请求后，Trailer可以在读取请求主体期间被修改， 
// 一旦请求主体返回EOF，调用者就不可再修改Trailer。 // 
// 很少有HTTP客户端、服务端或代理支持HTTP trailer。 
    
Trailer Header

// RemoteAddr允许HTTP服务器和其他软件记录该请求的来源地址，一般用于日志。 
// 本字段不是ReadRequest函数填写的，也没有定义格式。 
// 本包的HTTP服务器会在调用处理器之前设置RemoteAddr为"IP:port"格式的地址。 
// 客户端会忽略请求中的RemoteAddr字段。 
    RemoteAddr string

    // RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI // 
    // 一般应使用URI字段，在客户端设置请求的本字段会导致错误。
    RequestURI string

    // TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息 
    // 本字段不是ReadRequest函数填写的。 
    // 对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil。 
    // 客户端会忽略请求中的TLS字段。
    TLS *tls.ConnectionState

    Cancel <-chan struct{}

    Response *Response
    
}*/

// type URL struct {
//     Scheme     string
//     Opaque     string    // 编码不透明数据
//     User       *Userinfo // 用户名和密码信息
//     Host       string    // 主机或主机：端口
//     Path       string    // path (相对路径可以省略前导斜线)
//     RawPath    string    // 编码路径提示（请参阅EscapedPath方法）
//     ForceQuery bool      // 追加查询（'？'），即使RawQuery为空
//     RawQuery   string    // 编码的查询值，没有'？'
//     Fragment   string    // 引用的片段，不带'＃'
// }

func (h MyHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
    fmt.Println(*r)
    fmt.Println(r.Method)
    fmt.Println(*r.URL)
    fmt.Println(r.Host)
    fmt.Println(r.RemoteAddr)
    fmt.Println("UserAgent===>", r.UserAgent())
    fmt.Println("========================\r\n")
    fmt.Println(r.Header)
    fmt.Println("Accept======>",r.Header.Get("Accept"))

    fmt.Println("Cookie======>",r.Header.Get("Cookie"))

    fmt.Println("Referer=====>",r.Header.Get("Referer"))

   

    fmt.Println(r.Body)
    fmt.Println(r.GetBody)
    fmt.Println(r.URL.Path)

}

func main(){
    err := http.ListenAndServe(":9099",MyHandler{}); if err!=nil {
         fmt.Println("start http server fail:", err)
    }
}