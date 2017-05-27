#golang 内存管理
  内置运行时的编程语言通常会抛弃传统的内存分配方式， 改由自主管理。 这样可以完成类似预分配、内存池等操作， 以避开系统调用带来的性能问题。 当然， 还有一个重要的原因是为了更好地配合垃圾回收。
#基本策略
    1.每次从操作系统申请一大块内存（比如1MB）, 以减少系统调用；
    2.将申请到的大块内存按照特定的大小预先切分成小块， 构成链表；
    3.为对象分配内存时， 只需从大小合适的链表提取一个小块即可；
    4.回收对象内存时， 将该小块内存重新归还到原链表， 以便复用；
    5.如闲置内存过多， 则尝试归还部分内存给操作系统，降低整体开销。
    内存分配器只管理内存块， 并不关心对象状态。 且不会主动回收内存， 由垃圾回收器在完成清理操作后， 触发内存分配器回收操作。
#内存块
 ###分配器将其管理的内存块分为两种:
    span: 由多个地址连续的页（page）组成的大块内存；
    object: 将span 按特定大小切分成多个小块， 每个小块可存储一个对象。
 ###按照用途， span 面向内部管理， object 面向对象分配。
     分配器按页数来区分不同大小的span. 比如， 以页数为单位将span存放到管理数组中， 需要时就以页数为索引进行查找。 当然， span大小并非固定不变， 在获取闲置span时，如果没有找到大小合适的， 那就返回页数更多的， 此时会引发裁剪操作， 多余部分将构成新的span 被放回管理数组。 分配器还会尝试将地址相邻的空闲span 合并， 以构建更大的内存块， 减少碎片， 提供更灵活的分配策略。
#分配器
    优秀的内存分配器必须要在性能和内存利用率之间做到平衡。 好在， Golang 的起点很高， 直接采用了tcmalloc的成熟架构.

    分配器由三种组件组成。 
    - cache: 每个运行期工作线程都会绑定一个cache, 用于无锁object 分配。 
    - central: 为所有的cache 提供切分好的后备span次源。 
    - heap: 管理闲置span, 需要时向操作系统申请新内存。
#go1.7 源码中关于内存管理几个重要的文件
    1.runtime/malloc.go 这个文件里提供了make/new 等操作的分配内存函数， 是整个内存管理一个接口实现；
    2.runtime/mheap.go 从堆上分配内存， 管理页（4K）
    3.runtime/msize.go 为待分配的对象指定规格(size class)
    4.runtime/mfixalloc.go fixed-size 对象分配器。返回非0的内存。
    5.runtime/mcache.go 在Go中每个线程都有一个cache用于存储小对象.
    6.runtime/mstats.go 分配内存的统计
#几个重要的数据结
/ Main malloc heap.
// The heap itself is the "free[]" and "large" arrays,
// but all the other global data is here too.
type mheap struct {
    free      [_MaxMHeapList]mSpanList // free lists of given length
    freelarge mSpanList                // free lists length >= _MaxMHeapList
    busy      [_MaxMHeapList]mSpanList // busy lists of large objects of given length
    busylarge mSpanList                // busy lists of large objects length >= _MaxMHeapList
    allspans  **mspan                  // all spans out there
    gcspans   **mspan                  // copy of allspans referenced by gc marker or sweeper
    nspan     uint32
    sweepgen  uint32 // sweep generation, see comment in mspan
    sweepdone uint32 // all spans are swept
    // span lookup
    spans        **mspan
    spans_mapped uintptr
}

type fixalloc struct {  
    size   uintptr
    first  func(arg, p unsafe.Pointer) // called first time p is returned
    arg    unsafe.Pointer
    list   *mlink
    chunk  unsafe.Pointer
    nchunk uint32
    inuse  uintptr // in-use bytes now
    stat   *uint64
}

// mSpanList heads a linked list of spans.
//
// Linked list structure is based on BSD's "tail queue" data structure.
type mSpanList struct {  
    first *mspan  // first span in list, or nil if none
    last  **mspan // last span's next field, or first if none
}

type mspan struct {  
    next *mspan     // next span in list, or nil if none
    prev **mspan    // previous span's next field, or list head's first field if none
    list *mSpanList // For debugging. TODO: Remove.

}

type mcache struct {  
    // The following members are accessed on every malloc,
    // so they are grouped here for better caching.
    next_sample int32   // trigger heap sample after allocating this many bytes
    local_scan  uintptr // bytes of scannable heap allocated

    // Allocator cache for tiny objects w/o pointers.
    // See "Tiny allocator" comment in malloc.go.

    // tiny points to the beginning of the current tiny block, or
    // nil if there is no current tiny block.
    //
    // tiny is a heap pointer. Since mcache is in non-GC'd memory,
    // we handle it by clearing it in releaseAll during mark
    // termination.
    tiny             uintptr
    tinyoffset       uintptr
    local_tinyallocs uintptr // number of tiny allocs not counted in other stats

    // The rest is not accessed on every malloc.
    alloc [_NumSizeClasses]*mspan // spans to allocate from

    stackcache [_NumStackOrders]stackfreelist

    // Local allocator stats, flushed during GC.
    local_nlookup    uintptr                  // number of pointer lookups
    local_largefree  uintptr                  // bytes freed for large objects (>maxsmallsize)
    local_nlargefree uintptr                  // number of frees for large objects (>maxsmallsize)
    local_nsmallfree [_NumSizeClasses]uintptr // number of frees for small objects (<=maxsmallsize)
}

// Central list of free objects of a given size.
type mcentral struct {  
    lock      mutex
    sizeclass int32
    nonempty  mSpanList // list of spans with a free object, ie a nonempty free list
    empty     mSpanList // list of spans with no free objects (or cached in an mcache)
}

#小对象的分配流程
   1. 计算待分配对象对应规格(size class).
   2. 从cache.alloc 数组找到规格相同的span.
   3. 从span.freelist 链表提取可用的object.
   4. 如span.freelist 为空， 从central获取新的span.
   5. 如central.nonempty为空， 从heap.free/freelarge 获取， 并切分成object链表.
   6. 如heap 没有大小合适的闲置span, 操作系统申请新内存块。
#小对象的释放流程
    1.将标记为可回收object交还给所属span.freelist.
    2.该span 被放回central, 可供任意cache 重新获取使用.
    3.如span 已收回全部object, 则将其交还给heap, 以便重新切分复用；
    4.定期扫描heap里长时间闲置的span, 释放其占用内存。

作为工作线程私有且不被共享的cache 是实现高性能无锁分配的核心， 而central的作用是在多个cache间提高object的利用率， 避免内存浪费.

假如cache1 获取span后， 仅使用了一部分object, 那么剩余空间就可能会被浪费。 而回收操作将该span交还给central后， 该span完全可以被cache2, cacheN获取使用。 此时cache1已不在持有该span, 完全不会造 成问题。

将span 归还给heap, 是为了在不同规格object需求间平衡。

某时段某种规格的object需求量可能激增， 那么当需求过后， 大量被切分成该规格的span就会被闲置浪费。 将归还给heap， 就可以被其他需求获取， 重新切分。

#初始化
    因为内存分配器和垃圾回收算法都依赖连续的地址， 所以在初始化阶段， 预先保留了很大的一段虚拟地址空间。 保留地址空间并不会配内存。
    页所属的span指针区	 GC标记位图	         用户内存分配区域
    spans 512MB	      bitmap 32MB      	arena 512GB
    spans_mapped	  bitmap_mapped	    area_start	area_used	area_end

优化内存申请的方法--临时对象
带垃圾回收的语言， 虽然对于刚刚上手的程序员是友好的， 但是后期随着项目变得越来越巨大， 维护的内存问题也会逐渐暴露出来。 今天讲一种优化内存申请的方法--临时对象池。

在高并发的情况下， 如果每次请求都需要申请一块用于计算的内存， 比如：
    m := make([]int64, len(ids))
将会是一件成本很高的事， 为了定位项目中的慢语句， 我曾经采用"二分法”的方式打印语句， 定位程序变慢的代码位置。 它并不是每次都慢， 而是每过几秒钟就突然变得极其慢， TPS能从2000降到200。 引起这个问题就是类似于上面这条语句。 初始化一个slice, 初学者会用:
     m := make([]int64, 0)
高级一些的程序员都会知道， 这样第一次分配内存相当于没有分配， 如果要后续append元素， 会引起slice以指数形式扩充， 可以参考下面的代码，追加了3个元素， slice扩容了3次。
    a := make([]int64,  0)
    fmt.Println("cap=", cap(a), "len=", len(a))

    for i:=0; i<3; i++ {
        a = append(a, 1)
        fmt.Println("cap=", cap(a), "len=", len(a))
    }

cap= 0 len= 0  
cap= 1 len= 1  
cap= 2 len= 2  
cap= 4 len= 3  
每一次扩容空间， 都是会重新申请一块区域， 把就空间里面的元素复制进来， 把新的追加进来。 那旧空间里面的元素怎么办？ 等着垃圾回收呗。 简单的优化方式， 就是给自己要用的slice 提前申请好空间，类似于最开头的那行代码。
    make([]int64, 0, len(ids))
这样做了避免了多次扩容申请内存， 但还是有问题的。
堆还是栈？
程序会从操作系统申请一块内存， 而这块内存也会被分成堆和栈。 栈可以简单的理解成一次函数调用内部申请到的内存， 它们会随着函数的返回把内存还给系统。
    func F() {
        temp := make([]int, 0, 20)
        ...
    }
而上面这段代码， 申请的代码一模一样， 但是申请后作为返回值返回了， 编计器会认为变量之后还会被使用，当函数返回之后并不会将其内存归还， 那它就会被申请到堆上面了。 申请到堆上面的内存才会引起垃圾回收。

那么考考大家， 下面这三种情况怎么解释呢？
func F() {  
    a := make([]int, 0, 20)
    b := make([]int, 0, 20000)

    l := 20
    c := make([]int, 0, len(l))
}
a 和 b的代码一样，就是申请的空间不一样大， 但是它们两个的命运是截然相反的。 a前面已经介绍过了， 会申请到栈上， 而b由于申请内存较大， 编译器会把这种申请内存较大（>32KB）的变量转移到堆上面。 即使是临时变量， 申请过大也会在堆上面申请. 而c， 对我们而言， 其含义和a是一致的， 但是编译器对于这种不定长度的申请方式， 也会在堆上面申请， 即使申请的长度很短。

可以通过下面的命令查看变量申请的位置。 详细内容可以参考我之前的文章《【译】优化Go的模式》
    go build -gcflags="-m" . 2>&1
内存碎片化
实际项目基本都是通过c := make([]int, 0, l) 来申请内存， 长度都是不确定的。 自然而然这睦变量都会申请到堆上面了。 Golang使用的垃圾回收算法是【标记--清除】。 简单的说， 就是程序要从操作系统申请一块比较大的内存， 内存分成小块， 通过链表链接。 每次程序申请内存，就从链表上面遍历每一小块， 找到符合的就返回其地址， 没有合适的就从操作系统申请。 如果申请内存次数较我， 而且申请的大小不固定， 就会引起碎片化的问题。 申请的堆内存并没有用完， 但是用户申请的内存的时候却没有合适的空间提供。 这样会遍历整个链表， 还会继续向操作系统申请内存。 这样就能解释我一开始描述的问题， 申请一块内存变成了慢语句。
临时对象池
如何解决这个问题， 首先想到的就是对象池。 Golang 在sync 里面提供了对象池Pool。 一般大家都叫这个为对象池， 而我喜欢叫它临时对象池。 因为每次垃圾回收会把池子里面不被引用的对象回收掉。
func(p *Pool) Get() interface{}  
Get selects an arbitrary item from the Pool, removes it from the Pool, and >  
returns it to the caller. Get may choose to ignore the pool and treat it as  
empty. Callers should not assume any relation between values passed to Put and  
the values returned by Get.  
需要注意的是， Get方法会把返回的对象从池子里面删除。 所以用完了的对象， 还是得重新放回池子。 很快， 我写出了第一版对象池优化方案:
var idsPool = sync.Pool {  
    New: func() interface{} {
        ids := make([]int64, 0, 20000)
        return &ids
    },
}

func NewIds() []int64 {  
    ids := idsPool.Get().(*[]int64) 
    *ids = (*ids)[:0]
    idsPool.Put(ids)
    return *ids
}
这样的实现， 是把所有的slice都放到同一个池子里面了。 为了应对变长的问题， 都是按照一个较大的值申请的变量。 虽然是一种优化， 但是使用超大的slice计算， 性能并没有怎么提升。 紧接着参考了达达大神的代码sync_pool.go， 又写了一版:
var DEFAULT_SYNC_POOL *SyncPool

func NewPool() *SyncPool {  
    DEFAULT_SYNC_POOL = NewSyncPool(
        5, 30000, 2,
    )
    return DEFAULT_SYNC_POOL
}

func Alloc(size int) []int64 {  
    return DEFAULT_SYNC_POOL.Alloc(size)
}

func Free(mem []int64) {  
    DEFAULT_SYNC_POOL.Free(mem)
}

// SyncPool is a sync.Pool base slab allocation memory pool
type SyncPool struct {  
    classes     []sync.Pool
    classesSize []int
    minSize     int
    maxSize     int
}

func NewSyncPool(minSize, maxSize, factor int) *SyncPool {  
    n := 0
    for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
        n++
    }
    pool := &SyncPool{
        make([]sync.Pool, n),
        make([]int, n),
        minSize, maxSize,
    }
    n = 0
    for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
        pool.classesSize[n] = chunkSize
        pool.classes[n].New = func(size int) func() interface{} {
            return func() interface{} {
                buf := make([]int64, size)
                return &buf
            }
        }(chunkSize)
        n++
    }
    return pool
}

func (pool *SyncPool) Alloc(size int) []int64 {  
    if size <= pool.maxSize {
        for i := 0; i < len(pool.classesSize); i++ {
            if pool.classesSize[i] >= size {
                mem := pool.classes[i].Get().(*[]int64)
                // return (*mem)[:size]
                return (*mem)[:0]
            }
        }
    }
    return make([]int64, 0, size)
}

func (pool *SyncPool) Free(mem []int64) {  
    if size := cap(mem); size <= pool.maxSize {
        for i := 0; i < len(pool.classesSize); i++ {
            if pool.classesSize[i] >= size {
                pool.classes[i].Put(&mem)
                return
            }
        }
    }
}
调用例子:
attrFilters := cache.Alloc(len(ids))  
defer cache.Free(attrFilters)  
重点在Alloc方法。为了能支持变长的slice，这里申请了多个池子，其大小是从5开始，最大到30000，以2为倍数。也就是5、10、20……
DEFAULT_SYNC_POOL = NewSyncPool(  
    5,     
    30000, 
    2,     
)
分配内存的时候， 从池子里面找满足容量切最小的池子。 比如申请长度是2的， 就分配大小为5的那个池子。 如果是11， 就分配大小是20的那个池子里面的对象；
如果申请的slice很大， 超过了上限30000， 这种情况就不用使池子了， 直接从内存申请；
当然这些参数可以根据自己实际情况调整；
和之前的做法有所区别， 把对象重新放回池子是通Free方法实现的.
结论
为了优化接口， 前前后后搞了一年。 结果还是不错的。 TPS 提供了最小30%, TP99也降低了很多。
本篇文章来源 [http://blog.cyeam.com]
该博客还有介绍如下有趣的内容
go 工具链 http://blog.cyeam.com/golang/2016/09/27/go-tool-flags
- Golang 的垃圾回收 http://blog.cyeam.com/golang/2015/07/03/gogc
