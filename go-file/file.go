package main

import (
    "fmt"
   // "io/ioutil"
    "log"
    "net/http"
    "io"
    "strconv"
  //  "math"
    "os"
   // "strconv"
)

// 获取大小的借口
type Size interface {
    Size() int64
}

type Stat interface {
    Stat() (os.FileInfo, error)
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
    if "POST" == r.Method {
        name:= r.FormValue("name")
        
        total, _ := strconv.ParseInt(r.FormValue("total"), 10, 64)
        index, _ := strconv.ParseInt(r.FormValue("index"), 10, 64) 
         if total == index {
              fmt.Println("this is ok");
              return
         }

        fmt.Println(name,total,index)
        file, _, err := r.FormFile("data")
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        defer file.Close()
        fmt.Println(file)

         s:= make([]int64,total)
          fmt.Println("Capcity: ",cap(s),"Length: ",len(s))
         s = append(s,index)
         fmt.Println("Capcity: ",cap(s),"Length: ",len(s))
         count:= len(s) - 1
         fmt.Println("Count: ",count)
          fmt.Println(s)
        
        //  for i := 0; i < total; i++ {

        //  }
       
        // fileToBeChunked := "./somebigfile"

        // //file, err := os.Open(fileToBeChunked)

        // if err != nil {
        //     fmt.Println(err)
        //     os.Exit(1)
        // }

        // defer file.Close()

       // fileInfo, _ := file.Stat()

       // var fileSize int64 = fileInfo.Size()

        // const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement

        // // calculate total number of parts the file will be chunked into

        // totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

        // fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

        // for i := uint64(1); i < totalPartsNum; i++ {

        //     partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
        //     partBuffer := make([]byte, partSize)

        //     file.Read(partBuffer)

        //     // write to disk
        //     fileName := "somebigfile_" + strconv.FormatUint(i, 10)
        //     _, err := os.Create(fileName)

        //     if err != nil {
        //         fmt.Println(err)
        //         os.Exit(1)
        //     }

        //     // write/save buffer to disk
        //     ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

        //     fmt.Println("Split to : ", fileName)
        // }
    }
    // 上传页面
    w.Header().Add("Content-Type", "text/html")
    w.WriteHeader(200)
    html := `
     <script src="http://libs.baidu.com/jquery/1.11.1/jquery.min.js"></script>
     <script>
    var page = {
        init: function(){
            $("#upload").click($.proxy(this.upload, this));
        },
         
        upload: function(){
            var file = $("#file")[0].files[0],  //文件对象
                name = file.name,        //文件名
                size = file.size,        //总大小
                succeed = 0;
                  
            var shardSize = 2 * 1024 * 1024;     //以2MB为一个分片
            var shardCount = Math.ceil(size / shardSize);   //总片数
                  
            for(var i = 0;i < shardCount;i++){
                //计算每一片的起始与结束位置
                var start = i * shardSize;
                var end = Math.min(size, start + shardSize);
                var form = new FormData();//构造一个表单，FormData是HTML5新增的
                form.append("data", file.slice(start,end));  //slice方法用于切出文件的一部分
                form.append("name", name);
                form.append("total", shardCount);   //总片数
                form.append("index", i + 1);        //当前是第几片
                //Ajax提交
                $.ajax({
                    url: "/hello",
                    type: "POST",
                    data: form,
                    async: true,         //异步
                    processData: false,  //很重要，告诉jquery不要对form进行处理
                    contentType: false,  //很重要，指定为false才能形成正确的Content-Type
                    success: function(){
                        ++succeed;
                        $("#output").text(i + " / " + shardCount);
                    }
                });
            }
        }
    };
    $(function(){
        page.init();
    });
    </script>
    <input type="file" id="file" />
    <button id="upload">上传</button>
    <span id="output" style="font-size:12px">等待</span>
`
    io.WriteString(w, html)
}
func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}