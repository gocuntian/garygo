# docker swarm
##### http://dockone.io/article/168
## docker swarm
 Managers:

    -- swarm-00

 Workers:

    -- swarm-01

    -- swarm-02
    
    -- swarm-03 

## 创建swarm
#### docker swarm init --advertise-addr 172.20.14.132

 OR 

#### docker swarm init --listen-addr 172.20.14.132:2377

Swarm initialized: current node (8q2hlc0mlxnpuzvxn8yfmwlt7) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-27tlf9eh8ztrckdfie5ntv3f00xouzexaajktw4424ocpo6b4e-7cvlrbrc23wwlvs01o2iuaj1s 172.20.14.132:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
### 参数说明：

    --advertise-addr参数表示其它swarm中的worker节点使用此ip地址与manager联系。命令的输出包含了其它节点如何加入集群的命令。
使用docker info和docker node ls查看集群中的相关信息：

####  docker info
...
Swarm: active
NodeID: 41atspd62he1vshs4jmhpyufj
Is Manager: true
ClusterID: 5nl0kyz1dfmkgg2sx04vr8zoi
Managers: 1
Nodes: 1
Orchestration:
Task History Retention Limit: 5
...

#### docker node ls

ID                            HOSTNAME            STATUS              AVAILABILITY        MANAGER STATUS
8q2hlc0mlxnpuzvxn8yfmwlt7 *   xingcuntian         Ready               Active              Leader

## 将节点加入swarm

### swarm-01
#### docker swarm join --token SWMTKN-1-27tlf9eh8ztrckdfie5ntv3f00xouzexaajktw4424ocpo6b4e-7cvlrbrc23wwlvs01o2iuaj1s 172.20.14.132:2377

### swarm-02
#### docker swarm join --token SWMTKN-1-27tlf9eh8ztrckdfie5ntv3f00xouzexaajktw4424ocpo6b4e-7cvlrbrc23wwlvs01o2iuaj1s 172.20.14.132:2377

### swarm-03
#### docker swarm join --token SWMTKN-1-27tlf9eh8ztrckdfie5ntv3f00xouzexaajktw4424ocpo6b4e-7cvlrbrc23wwlvs01o2iuaj1s 172.20.14.132:2377

#### 注意：如果当时创建swarm时没有记下命令的输出，可以通过在manager节点上运行docker swarm join-token worker命令来获取如何加入swarm的命令。

## 在Swarm中部署服务
### swarm-00 manager节点上运行如下命令来部署服务： 
#### docker service create --replicas 1 --name ping00 alpine ping docker.com

OR

#### docker service create --name ping00 alpine ping docker-swarm-00
### 参数说明：

    --replicas参数指定服务由几个实例组成。最后的命令行参数alpine ping docker.com指定了使用alpine镜像创建服务，实例启动时运行ping docker.com命令。这与docker run命令是一样的。
使用docker service ls查看正在运行服务的列表：

#### docker service ls
#### docker service tasks ping00


### 查询Swarm中服务的信息

    在部署了服务之后，登录到manager节点，运行下面的命令来显示服务的信息。参数--pretty使命令输出格式化为可读的格式，不加--pretty可以输出更详细的信息：

#### docker service inspect --pretty ping00

    使用命令docker service ps <SERVICE-ID>可以查询到哪个节点正在运行该服务：

#### docker service ps ping00

### swarm-01
#### docker logs -f CONTAINER ID

## 在Swarm中动态扩展服务

登录到manager节点，使用命令docker service scale <SERVICE-ID>=<NUMBER-OF-TASKS>来将服务扩展到指定的实例数：

#### docker service scale ping00=3


## 删除Swarm中的服务

    在manager节点上运行docker service rm ping00 便可以将服务删除。删除服务时，会将服务在各个节点上创建的容器一同删除，而并不是将容器停止。

此外Swarm模式还提供了服务的滚动升级，将某个worker置为维护模式，及路由网等功能。在Docker将Swarm集成进Docker引擎后，可以使用原生的Docker CLI对容器集群进行各种操作，使集群的部署更加方便、快捷。



