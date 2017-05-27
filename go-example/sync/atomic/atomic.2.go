package main

// 创建一个文件存放数据,在同一时刻,可能会有多个Goroutine分别进行对此文件的写操作和读操作.
// 每一次写操作都应该向这个文件写入若干个字节的数据,作为一个独立的数据块存在,这意味着写操作之间不能彼此干扰,写入的内容之间也不能出现穿插和混淆的情况
// 每一次读操作都应该从这个文件中读取一个独立完整的数据块.它们读取的数据块不能重复,且需要按顺序读取.
// 例如: 第一个读操作读取了数据块1,第二个操作就应该读取数据块2,第三个读操作则应该读取数据块3,以此类推
// 对于这些读操作是否可以被同时执行,不做要求. 即使同时进行,也应该保持先后顺序.
import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Data []byte

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	Rsn() int64
	Wsn() int64
	DataLen() uint32
}

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	rcond   *sync.Cond //读操作需要用到的条件变量
	woffset int64
	roffset int64
	wmutex  sync.Mutex
	rmutex  sync.Mutex
	dataLen uint32
}

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Fail to find", f, "cServer start Failed")
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}

	df := &myDataFile{
		f:       f,
		dataLen: dataLen,
	}
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	// 读互斥锁定
	df.rmutex.Lock()
	offset = df.roffset
	// 更改偏移量, 当前偏移量+数据块长度
	df.roffset += int64(df.dataLen)
	// 读互斥解锁
	df.rmutex.Unlock()

	//读取一个数据块,最后读取的数据块序列号
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	//读写锁:读锁定
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				//暂时放弃fmutex的 读锁,并等待通知的到来
				df.rcond.Wait()
				continue
			}
		}
		break
	}
	d = bytes
	return

}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	//读取并更新写的偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块,最后写入数据块的序号
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	//发送通知
	df.rcond.Signal()
	return
}

func (df *myDataFile) Rsn() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) Wsn() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func main() {
	var dataFile DataFile
	dataFile, _ = NewDataFile("./mutex.dat", 10)

	var d = map[int]Data{
		1: []byte("xingcuntian1"),
		2: []byte("xingcuntian2"),
		3: []byte("xingcuntian3"),
	}

	for i := 1; i < 4; i++ {
		go func(i int) {
			wsn, _ := dataFile.Write(d[i])
			fmt.Println("write i=", i, ",wsn=", wsn, ",success.")
		}(i)
	}

	for i := 1; i < 4; i++ {
		go func(i int) {
			rsn, d, _ := dataFile.Read()
			fmt.Println("Read i=", i, ",rsn=", rsn, ",data=", d, ",success.")
		}(i)
	}

	time.Sleep(10 * time.Second)
}

// func (f *File) ReadAt(b []byte, off int64) (n int, err error)
// ReadAt从指定的位置（相对于文件开始位置）读取len(b)字节数据并写入b。它返回读取的字节数和可能遇到的任何错误。当n<len(b)时，本方法总是会返回错误；如果是因为到达文件结尾，返回值err会是io.EOF。
