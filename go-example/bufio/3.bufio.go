package main

import (
    "fmt"
    "bufio"
    "os"
)
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// Fprintf根据format参数生成格式化的字符串并写入w。返回写入的字节数和遇到的任何错误。
// func (b *Writer) Flush() error
// Flush方法将缓冲中的数据写入下层的io.Writer接口。
func main(){
    fmt.Fprintf(os.Stdout,"%s\n","hello world! - unbuffered")
    buf :=bufio.NewWriter(os.Stdout)

    fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
    buf.Flush()
}
// fmt.Fprintf() 函数的实际签名

// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) 
// 其不是写入一个文件，而是写入一个 io.Writer 接口类型的变量，下面是 Writer 接口在 io 包中的定义：

// type Writer interface {
//     Write(p []byte) (n int, err error)
// }
// fmt.Fprintf() 依据指定的格式向第一个参数内写入字符串，第一参数必须实现了 io.Writer 接口。Fprintf() 能够写入任何类型，只要其实现了 Write 方法，包括 os.Stdout,文件（例如 os.File），管道，网络连接，通道等等，同样的也可以使用 bufio 包中缓冲写入。

// bufio 包中定义了 type Writer struct{...}

// bufio.Writer 实现了 Write 方法：

// func (b *Writer) Write(p []byte) (nn int, err error)
// 它还有一个工厂函数：传给它一个 io.Writer 类型的参数，它会返回一个缓冲的 bufio.Writer 类型的 io.Writer:

// func NewWriter(wr io.Writer) (b *Writer)
// 其适合任何形式的缓冲写入。

// 在缓冲写入的最后千万不要忘了使用 Flush()，否则最后的输出不会被写入。