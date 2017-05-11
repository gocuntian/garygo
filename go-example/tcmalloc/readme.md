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
