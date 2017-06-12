go使用etcd3构建微服务发现和配置管理
今天Go 1.7正式版发布了！

etcd是一个高可用的键值存储系统，主要用于共享配置和服务发现。etcd是由CoreOS开发并维护的，灵感来自于 ZooKeeper 和 Doozer，它使用Go语言编写，并通过Raft一致性算法处理日志复制以保证强一致性。

go-etcd 这个已经被废弃，建议使用官方的 github.com/coreos/etcd/clientv3

安装&运行etcd

brew install etcd

etcd

使用etcdctl命令

etcdctl 是一个命令行客户端，它能提供一些简洁的命令，供用户直接跟 etcd 服务打交道，而无需基于 HTTP API 方式。这在某些情况下将很方便，例如对服务进行测试或者手动修改数据库内容。也推荐在刚接触 etcd 时通过 etcdctl 命令来熟悉相关操作，这些操作跟 HTTP API 实际上是对应的。

安装 go get github.com/coreos/etcd/etcdctl 也可以直接下载etcd二进制 （包含etcd、etcdctl）

注意：目前本文测试使用的是etcd v3，一定不要忘记在环境变量中设置 export ETCDCTL_API=3 否则etcdctl默认使用的是v2与v3是完全隔离的（同一个etcd服务，不同的存储引擎）。

API V2与V3区别

事务：ETCD V3提供了多键条件事务（multi-key conditional transactions），应用各种需要使用事务代替原来的Compare-And-Swap操作。
平键空间（Flat key space）：ETCD V3不再使用目录结构，只保留键。例如：”/a/b/c/“是一个键，而不是目录。V3中提供了前缀查询，来获取符合前缀条件的所有键值，这变向实现了V2中查询一个目录下所有子目录和节点的功能。
简洁的响应：像DELETE这类操作成功后将不再返回操作前的值。如果希望获得删除前的值，可以使用事务，来实现一个原子操作，先获取键值，然后再删除。
租约：租约代替了V2中的TTL实现，TTL绑定到一个租约上，键再附加到这个租约上。当TTL过期时，租约将被销毁，同时附加到这个租约上的键也被删除。


存储 etcdctl put sensors "{aa:1, bb: 2}"

获取 etcdctl get sensors

监视 etcdctl watch sensors

获取所有值（或指定前缀 ）etcdctl get --prefix=true ""