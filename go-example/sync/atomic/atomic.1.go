package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// 创建一个文件存放数据,在同一时刻,可能会有多个Goroutine分别进行对此文件的写操作和读操作.
// 每一次写操作都应该向这个文件写入若干个字节的数据,作为一个独立的数据块存在,这意味着写操作之间不能彼此干扰,写入的内容之间也不能出现穿插和混淆的情况
// 每一次读操作都应该从这个文件中读取一个独立完整的数据块.它们读取的数据块不能重复,且需要按顺序读取.
// 例如: 第一个读操作读取了数据块1,第二个操作就应该读取数据块2,第三个读操作则应该读取数据块3,以此类推
// 对于这些读操作是否可以被同时执行,不做要求. 即使同时进行,也应该保持先后顺序.

type Data []byte

type DataFile interface {
	// 读取一个数据块
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号
	Rsn() int64
	//获取最后写入的数据块的序列号
	Wsn() int64
	//获取数据块的长度
	DataLen() uint32
}

//数据文件的实现类型
type myDataFile struct {
	f       *os.File     //文件
	fmutex  sync.RWMutex //被用于文件的读写锁
	rcond   *sync.Cond   //读操作需要用到的条件变量
	woffset int64        //写操作需要用到的偏移量
	roffset int64        //读操作需要用到的偏移量
	dataLen uint32       //数据块长度
}

//初始化DataFile类型值的函数，返回一个DataFile类型的值
//初始化DataFile类型值的函数,返回一个DataFile类型的值
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	//f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
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
	//创建一个可用的条件变量(初始化),返回一个*sync.Cond类型的结果值,我们就可以调用该值拥有的三个方法Wait,Signal,Broadcast
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

//获取并更新读偏移量,根据读偏移量从文件中读取一块数据,把该数据块封装成一个Data类型值并将其作为结果值返回
func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	for {
		//LoadInt64原子性的获取*addr的值。
		offset = atomic.LoadInt64(&df.roffset)
		//CompareAndSwapInt32原子性的比较*addr和old，如果相同则将new赋值给*addr并返回真。
		if atomic.CompareAndSwapInt64(&df.roffset, offset, (offset + int64(df.dataLen))) {
			break
		}
	}
	// 读取一个数据块,最后读取的数据块序列号
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
	for {
		offset = atomic.LoadInt64(&df.woffset)
		if atomic.CompareAndSwapInt64(&df.woffset, offset, (offset + int64(df.dataLen))) {
			break
		}
	}
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
	////发送通知
	df.rcond.Signal()
	return
}

func (df *myDataFile) Rsn() int64 {
	offset := atomic.LoadInt64(&df.roffset)
	return offset / int64(df.dataLen)
}

func (df *myDataFile) Wsn() int64 {
	offset := atomic.LoadInt64(&df.woffset)
	return offset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func main() {
	var dataFile DataFile
	dataFile, _ = NewDataFile("./mutex_2017.dat", 10)
	var d = map[int]Data{
		1: []byte("xingcuntian_1"),
		2: []byte("xingcuntian_2"),
		3: []byte("xingcuntian_3"),
	}
	//写入数据
	for i := 1; i < 4; i++ {
		go func(i int) {
			wsn, _ := dataFile.Write(d[i])
			fmt.Println("write i = ", i, ", wsn=", wsn, ",success.")
		}(i)
	}

	//读取数据
	for i := 1; i < 4; i++ {
		go func(i int) {
			rsn, d, _ := dataFile.Read()
			fmt.Println("Read i=", i, ",rsn=", rsn, ",data=", d, ",success.")
		}(i)
	}
	time.Sleep(3 * time.Second)
}

// type Cond struct {
//     // 在观测或更改条件时L会冻结
//     L Locker
//     // 包含隐藏或非导出字段
// }
// Cond实现了一个条件变量，一个线程集合地，供线程等待或者宣布某事件的发生。
// 每个Cond实例都有一个相关的锁（一般是*Mutex或*RWMutex类型的值），它必须在改变条件时或者调用Wait方法时保持锁定。Cond可以创建为其他结构体的字段，Cond在开始使用后不能被拷贝。

// func NewCond(l Locker) *Cond
// 使用锁l创建一个*Cond。

// func (c *Cond) Broadcast()
// Broadcast唤醒所有等待c的线程。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定。

// func (c *Cond) Signal()
// Signal唤醒等待c的一个线程（如果存在）。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定。

// func (c *Cond) Wait()
// Wait自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L。和其他系统不同，Wait除非被Broadcast或者Signal唤醒，不会主动返回。
// 因为线程中Wait方法是第一个恢复执行的，而此时c.L未加锁。调用者不应假设Wait恢复时条件已满足，相反，调用者应在循环中等待：
