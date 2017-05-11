package codec

import (
    "bufio"
    "bytes"
    "encoding/binary"
)
// 解决粘包问题有多种多样的方式, 我们这里的做法是:

// 发送方在每次发送消息时将消息长度写入一个int32作为包头一并发送出去, 我们称之为Encode
// 接受方则先读取一个int32的长度的消息长度信息, 再根据长度读取相应长的byte数据, 称之为Decode
// |------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// type Buffer struct {
//     // 内含隐藏或非导出字段
// }
// Buffer是一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲。

// func NewBuffer(buf []byte) *Buffer
// NewBuffer使用buf作为初始内容创建并初始化一个Buffer。本函数用于创建一个用于读取已存在数据的buffer；也用于指定用于写入的内部缓冲的大小，此时，buf应为一个具有指定容量但长度为0的切片。buf会被作为返回值的底层缓冲切片。

// func (b *Buffer) Bytes() []byte
// 返回未读取部分字节数据的切片，len(b.Bytes()) == b.Len()。如果中间没有调用其他方法，修改返回的切片的内容会直接改变Buffer的内容。
// |-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

// var LittleEndian littleEndian
// 小端字节序的实现。

// func Write(w io.Writer, order ByteOrder, data interface{}) error
// 将data的binary编码格式写入w，data必须是定长值、定长值的切片、定长值的指针。order指定写入数据的字节序，写入结构体时，名字中有'_'的字段会置为0。

// func Read(r io.Reader, order ByteOrder, data interface{}) error
// 从r中读取binary编码的数据并赋给data，data必须是一个指向定长值的指针或者定长值的切片。从r读取的字节使用order指定的字节序解码并写入data的字段里当写入结构体是，名字中有'_'的字段会被跳过，这些字段可用于填充（内存空间）。
// |---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

// func (b *Reader) Buffered() int
// Buffered返回缓冲中现有的可读取的字节数。

// func (b *Reader) Peek(n int) ([]byte, error)
// Peek返回输入流的下n个字节，而不会移动读取位置。返回的[]byte只在下一次调用读取操作前合法。如果Peek返回的切片长度比n小，它也会返会一个错误说明原因。如果n比缓冲尺寸还大，返回的错误将是ErrBufferFull。
// |----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

func Encode(message string) ([]byte, error){
    //读取消息的长度
    var length int32 = int32(len(message))
    var pkg *bytes.Buffer = new(bytes.Buffer)
    //写入消息头
    err := binary.Write(pkg,binary.LittleEndian,length)
    if err != nil {
        return nil, err
    }

    //写入消息实体
    err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
    if err != nil {
        return nil, err
    }

    return pkg.Bytes(), nil

}

func Decode(reader *bufio.Reader) (string, error){
    //读取消息的长度
    lengthByte, _ := reader.Peek(4)
    lengthBuff := bytes.NewBuffer(lengthByte)
    var length int32
    err := binary.Read(lengthBuff,binary.LittleEndian,&length)
    if err != nil {
        return "", err
    }
    if int32(reader.Buffered()) < length+4 {
        return "", err
    }

    //读取消息真正的内容
    pack := make([]byte,int(4+length))
    _, err = reader.Read(pack)
    if err != nil {
        return "", err
    }
    return string(pack[4:]), nil
}