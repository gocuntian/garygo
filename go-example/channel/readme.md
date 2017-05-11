Go的Goroutine能够让我们并行的运行一些代码。但是要有效的利用它，需要一些额外的工作。当进程完成创建，我们应该能够将数据传递到正在运行的进程，我们也应该能够获取数据从正在运行的进程，channels 做到了这一点并且能够很好的与goroutines工作。

我们可以把Channel想象成一个通道管或定义的大小和容量的传输带，一边可以往管道上面存放东西，另外一边可以将其取出

Channels用来同步并发执行的函数并提供它们某种传值交流的机制。Channels的一些特性：通过channel传递的元素类型、容器（或缓冲区）和传递的方向由“<-”操作符指定。你可以使用内置函数 make分配一个channel


i := make(chan int)       // by default the capacity is 0
s := make(chan string, 3) // non-zero capacity
 
r := make(<-chan bool)          // can only read from
w := make(chan<- []os.FileInfo) // can only write to

通过<- 操作符发送或接收的数据


my_channel := make(chan int)

//within some goroutine - to put a value on the channel
my_channel <- 5 

//within some other goroutine - to take a value off the channel
var my_recvd_value int
my_recvd_value = <- my_channel

my_channel := make(chan int)
 
//within some goroutine - to put a value on the channel
my_channel <- 5 
 
//within some other goroutine - to take a value off the channel
var my_recvd_value int
my_recvd_value = <- my_channel

ic_send_only := make (<-chan int) //a channel that can only send data - arrow going out is sending
ic_recv_only := make (chan<- int) //a channel that can only receive a data - arrow going in is receiving


接收Channle的数据的时候我们会遇到，什么时候需要停止等待数据。有更多数据，还是已经全部完成？我们是继续等待还是继续？一种方法就是不断的轮询和检查通道是否已经关闭，但是这种方法并不是特别有效。

Channels and range
Go提供了range关键词,当它与Channel 一起使用的时候他会等待channel的关闭。

Channels and select
golang 的 select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。注意到 select 的代码形式和 switch 非常相似， 不过 select 的 case 里的操作语句只能是  IO 操作