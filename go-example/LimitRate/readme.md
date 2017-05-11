Go 提供了一个package(golang.org/x/time/rate) 用来方便的对速度进行限制,下面就用示例来说明下如何使用。

首先创建一个rate.Limiter,其有两个参数，第一个参数为每秒发生多少次事件，第二个参数是其缓存最大可存多少个事件。

rate.Limiter提供了三类方法用来限速

Wait/WaitN 当没有可用或足够的事件时，将阻塞等待 推荐实际程序中使用这个方法
Allow/AllowN 当没有可用或足够的事件时，返回false
Reserve/ReserveN 当没有可用或足够的事件时，返回 Reservation，和要等待多久才能获得足够的事件。