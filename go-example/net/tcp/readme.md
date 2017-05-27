#长连接的TCP服务器与客户端通讯有个基本的认识:
### client向server发起连接，server接受client连接，双方建立连接。Client与server完成一次读写之后，它们之间的连接并不会主动关闭，后续的读写操作会继续使用这个连接

#服务器的粘包处理
### 什么是粘包

一个完整的消息可能会被TCP拆分成多个包进行发送，也有可能把多个小的包封装成一个大的数据包发送，这个就是TCP的拆包和封包问题

### TCP粘包和拆包产生的原因

应用程序写入数据的字节大小大于套接字发送缓冲区的大小

进行MSS大小的TCP分段。MSS是最大报文段长度的缩写。MSS是TCP报文段中的数据字段的最大长度。数据字段加上TCP首部才等于整个的TCP报文段。所以MSS并不是TCP报文段的最大长度，而是：MSS=TCP报文段长度-TCP首部长度

以太网的payload大于MTU进行IP分片。MTU指：一种通信协议的某一层上面所能通过的最大数据包大小。如果IP层有一个数据包要传，而且数据的长度比链路层的MTU大，那么IP层就会进行分片，把数据包分成托干片，让每一片都不超过MTU。注意，IP分片可以发生在原始发送端主机上，也可以发生在中间路由器上。

### TCP粘包和拆包的解决策略

消息定长。例如100字节。
在包尾部增加回车或者空格符等特殊字符进行分割，典型的如FTP协议
将消息分为消息头和消息尾。
其它复杂的协议，如RTMP协议等。
参考文章 (http://blog.csdn.net/initphp/article/details/41948919)

## 处理方式

解决粘包问题有多种多样的方式, 这里的做法是:

发送方在每次发送消息时将消息长度写入一个int32作为包头一并发送出去, 我们称之为Encode
接受方则先读取一个int32的长度的消息长度信息, 再根据长度读取相应长的byte数据, 称之为Decode