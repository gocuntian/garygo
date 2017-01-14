package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "bufio"
)

func check(e error){
    if e!=nil{
        panic(e)
    }
}

func main(){
    //func WriteFile(filename string, data []byte, perm os.FileMode) error
    //函数向filename指定的文件中写入数据。
    //如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
    d1:=[]byte("hello\ngos\n")
    err:=ioutil.WriteFile("/tmp/data1",d1,0644)
    check(err)

    //func Create(name string) (file *File, err error)
    //Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，
    //如果文件已存在会截断它（为空文件）。如果成功，返回的文件对象可用于I/O；
    //对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。
    f,err:=os.Create("/tmp/data2")
    check(err)
    //func (f *File) Close() error
    //Close关闭文件f，使文件不能用于读写。它返回可能出现的错误。
    defer f.Close()

    d2 := []byte{115,111,109,101,10}
    //func (f *File) Write(b []byte) (n int, err error)
    //Write向文件中写入len(b)字节数据。
    //它返回写入的字节数和可能遇到的任何错误。
    //如果返回值n!=len(b)，本方法会返回一个非nil的错误。
    n2,err:=f.Write(d2)
    check(err)
    fmt.Printf("write %d bytes\n",n2)
    //func (f *File) WriteString(s string) (ret int, err error)
    //WriteString类似Write，但接受一个字符串参数。
    n3,err:=f.WriteString("writes\n")
    fmt.Printf("write %d bytes\n",n3)
    //func (f *File) Sync() (err error)
    //Sync递交文件的当前内容进行稳定的存储。一般来说，
    //这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存。
    f.Sync()

    //bufio
    //func NewWriter(w io.Writer) *Writer
    //NewWriter创建一个具有默认大小缓冲、写入w的*Writer。
    w:=bufio.NewWriter(f)
    //func (b *Writer) WriteString(s string) (int, error)
    //WriteString写入一个字符串。返回写入的字节数。如果返回值nn < len(s)，
    //还会返回一个错误说明原因。
    n4,err:=w.WriteString("buffered\n")
    fmt.Printf("write %d bytes\n",n4)
    //func (b *Writer) Flush() error
    //Flush方法将缓冲中的数据写入下层的io.Writer接口。
    w.Flush()

}