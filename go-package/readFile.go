package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error){
    if e!=nil{
        panic(e)
    }
}

func main(){
    //ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。
    dat, err:=ioutil.ReadFile("/tmp/xct.txt")
    check(err)
    fmt.Print(string(dat))
    //Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
    f,err:=os.Open("/tmp/xct.txt")
    check(err)
    b1:=make([]byte,10)
    //Read方法从f中读取最多len(b)字节数据并写入b。它返回读取的字节数和可能遇到的任何错误。文件终止标志是读取0个字节且返回值err为io.EOF。
    n1,err:=f.Read(b1)
    check(err)
    fmt.Printf("%d bytes:%s\n",n1,string(b1))
    //Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。
    o2,err:=f.Seek(6,0)
    check(err)
    b2:=make([]byte,2)
    n2,err:=f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2))
     //func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
    //ReadAtLeast从r至少读取min字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。 如果min比buf的长度还大，函数会返回ErrShortBuffer。只有返回值err为nil时，返回值n才会不小于min
    o3,err:=f.Seek(6,0)
    check(err)
    b3:=make([]byte,2)
    n3,err:=io.ReadAtLeast(f,b3,2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3))

    //func NewReader(rd io.Reader) *Reader
    //NewReader创建一个具有默认大小缓冲、从r读取的*Reader。
    _,err=f.Seek(0,0)
    check(err)
    r4:=bufio.NewReader(f)


    //func (b *Reader) Peek(n int) ([]byte, error)
    //Peek返回输入流的下n个字节，而不会移动读取位置。
    //返回的[]byte只在下一次调用读取操作前合法。
    //如果Peek返回的切片长度比n小，
    //它也会返会一个错误说明原因。如果n比缓冲尺寸还大，
    //返回的错误将是ErrBufferFull。
    b4,err:=r4.Peek(9)
    check(err)
    fmt.Printf("5 bytes: %s\n",string(b4))
    f.Close()



}