Channel机制：

相对sync.WaitGroup而言，golang中利用channel实现同步则简单的多．channel自身可以实现阻塞，其通过<-进行数据传递，channel是golang中一种内置基本类型，对于channel操作只有４种方式：

创建channel(通过make()函数实现，包括无缓存channel和有缓存channel);

向channel中添加数据（channel<-data）;

从channel中读取数据（data<-channel）;

关闭channel(通过close()函数实现，关闭之后无法再向channel中存数据，但是可以继续从channel中读取数据）


channel分为有缓冲channel和无缓冲channel,两种channel的创建方法如下:

var ch = make(chan int) //无缓冲channel,等同于make(chan int ,0)

var ch = make(chan int,5) //有缓冲channel,缓冲大小是５

其中无缓冲channel在读和写是都会阻塞，而有缓冲channel在向channel中存入数据没有达到channel缓存总数时，可以一直向里面存，直到缓存已满才阻塞．由于阻塞的存在，所以使用channel时特别注意使用方法，防止死锁的产生




对于无缓存的channel,放入channel和从channel中向外面取数据这两个操作不能放在同一个协程中，防止死锁的发生；同时应该先利用go 开一个协程对channel进行操作，此时阻塞该go 协程，然后再在主协程中进行channel的相反操作（与go 协程对channel进行相反的操作），实现go 协程解锁．即必须go协程在前，解锁协程在后．

带缓存channel: 
对于带缓存channel，只要channel中缓存不满，则可以一直向 channel中存入数据，直到缓存已满；同理只要channel中缓存不为０，便可以一直从channel中向外取数据，直到channel缓存变为０才会阻塞． 

由此可见，相对于不带缓存channel，带缓存channel不易造成死锁，可以同时在一个goroutine中放心使用，

golang 并发总结： 
并发两种方式：sync.WaitGroup，该方法最大优点是Wait()可以阻塞到队列中的所有任务都执行完才解除阻塞，但是它的缺点是不能够指定并发协程数量． 
channel优点：能够利用带缓存的channel指定并发协程goroutine，比较灵活．但是它的缺点是如果使用不当容易造成死锁；并且他还需要自己判定并发goroutine是否执行完． 

但是相对而言，channel更加灵活，使用更加方便，同时通过超时处理机制可以很好的避免channel造成的程序死锁，因此利用channel实现程序并发，更加方便，更加易用． 


Go的临时对象池sync.Pool
在高并发或者大量的数据请求的场景中，我们会遇到很多问题，垃圾回收就是其中之一（garbage collection），为了减少优化GC，我们一般想到的方法就是能够让对象得以重用。这就需要一个对象池来存储待回收对象，等待下次重用，从而减少对象产生数量。我们可以把sync.Pool类型值看作是存放可被重复使用的值的容器。此类容器是自动伸缩的、高效的，同时也是并发安全的。为了描述方便，我们也会把sync.Pool类型的值称为临时对象池，而把存于其中的值称为对象值。这个类设计的目的是用来保存和复用临时对象，以减少内存分配，降低CG压力。

sync.Pool 最常用的两个函数Get/Put

var pool = &sync.Pool{New:func()interface{}{return NewObject()}}
    pool.Put()
    Pool.Get()
对象池在Get的时候没有里面没有对象会返回nil，
所以我们需要New function来确保当获取对象对象池为空时，重新生成一个对象返回，
前者的功能是从池中获取一个interface{}类型的值，
而后者的作用则是把一个interface{}类型的值放置于池中。