现在 Go 1.9 beta版已发布， 正式版预期在8月初发布，让我们先来看看你Go 1.9带来了那些新特性。

type alias

    类型别名原本在1.8中加入的，但是临时发现有些问题，为了能全面的设计 type alias 被移到了 Go 1.9中了。

    这个特性主要用在类型从一个package移动到另外一个package中的时候，导致的项目中对引入的路径不一致导致的问题， 比如原先 context是在golang.org／x/net/context包下，在Go 1.7中菜正式移到标准库context。

    相关的issue: go#16339 go#18130
    提案: type alias

并发map

    在Go 1.6之前， 内置的map类型是部分goroutine安全的，并发的读没有问题，并发的写可能有问题。自go 1.6之后， 并发地读写map会报错，这在一些知名的开源库中都存在这个问题，所以go 1.9之前的解决方案是额外绑定一个锁，封装成一个新的struct或者单独使用锁都可以。

    群众的呼声是响亮的，并发map在项目中大量使用，所以Go 1.9中在包sync加入了新的map， 查询、存储和删除都是平均常数时间，可以并发访问。

    Monotonic Time
    先前的time包的实现都是基于wall time的，但是当机器的时钟调整后会有问题。 比如在计算duration的时候，如果时钟往回拨，可能导致end时间比start时间还早。
    所以Go 1.9使用monotonic Time来实现大部分的time中的函数，在计算duration的时候不会出现因为时钟调整出现的误差了。
    设计文档: monotonic time
    位处理操作

新增加了math/bits包， 提供了很多位运算的函数。

    Test Helper函数

    新加｀(T).Helper和(B).Helper m｀， 用来标记调用的函数是一个测试辅助函数，当输出文件名和行数的时候，这个函数回呗忽略。

标准库的微小改动

标准库也有一些小的功能的加入和提升， 比如image、net、runtime、sync等。

并行编译

支持并行地编译函数，并且在Go 1.9中势默认设置。如果不想并行编译，设置GO19CONCURRENTCOMPILATION为0。

./... 会忽略vendor下的包

这一条很有用，以后你在Makefile中可以直接使用./...,而不是曲折地将vendor文件夹排除。

如果你想使用vendor下的包， 可以使用./vendor/...通配符。

性能提升

性能提升多少势很难精确描述的，对于大部分的程序，应该运行更快一点。

主要在于垃圾回收器的优化、更好的生成的代码以及核心库的优化。