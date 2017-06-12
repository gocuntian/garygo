
启动首个节点
consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=a1 -bind=192.168.100.104 -config-dir /etc/consul.d 
consul members
加入集群
consul join 192.168.100.103


netstat  -tunpea | grep consul


ARG	Comment
-server	以服务模式运行
-bootstrap-expect	指定期望加入的节点数
-data-dir	指定数据存放的位置
-node	指定节点名
-bind	指定绑定的IP
-config-dir	指定配置目录
-ui  UI服务
-ui-dir  

启动第二个节点

consul agent -data-dir /tmp/consul -node=a2 -bind=192.168.100.103 -config-dir /etc/consul.d

consul members

查询节点
可以通过 DNS API 或 HTTP API 来查询节点

如果使用DNS API，查询结构为 NAME.node.consul 和 NAME.node.DATACENTER.consul
104
dig @127.0.0.1 -p 8600 a2.node.consul
dig @127.0.0.1 -p 8600 a2.node.dc1.consul

dig @127.0.0.1 -p 8600 a1.node.dc1.consul


