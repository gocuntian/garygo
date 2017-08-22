# Docker Swarm
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

## demo 使用

#### docker service create --name wesite --publish 80:80 sixeyed/docker-swarm-walkthrough

#  Docker Machine

    https://docs.docker.com/machine/install-machine/

    curl -L https://github.com/docker/machine/releases/download/v0.12.2/docker-machine-`uname -s`-`uname -m` >/tmp/docker-machine &&

    chmod +x /tmp/docker-machine && sudo cp /tmp/docker-machine /usr/local/bin/docker-machine

#### docker-machine version
#### docker-machine ls
#### docker-machine create --driver virtualbox default
#### docker-machine ls
 NAME      ACTIVE   DRIVER       STATE     URL                         SWARM   DOCKER   ERRORS
 default   *        virtualbox   Running   tcp://192.168.99.187:2376           v1.9.1
#### docker-machine env default
 export DOCKER_TLS_VERIFY="1"
 export DOCKER_HOST="tcp://172.16.62.130:2376"
 export DOCKER_CERT_PATH="/Users/<yourusername>/.docker/machine/machines/default"
 export DOCKER_MACHINE_NAME="default"
#### eval "$(docker-machine env default)"
#### docker-machine ip default

#### docker run -d -p 8000:80 nginx

#### curl $(docker-machine ip default):8000 
#### docker-machine stop default
#### docker-machine start default


 # Run this command to configure your shell:
 # eval "$(docker-machine env default)"

#### docker-machine rm <machine-name>
#### docker-machine rm -f $(docker-machine ls -q)




#### $ docker-machine ls

 NAME      ACTIVE   URL          STATE     URL                         SWARM   DOCKER    ERRORS
 default   *       virtualbox   Running   tcp://192.168.99.100:2376           v1.9.1

#### $ docker run swarm --help




# Demo
## Docker Swarm Cluster using Consul

https://www.ibm.com/developerworks/cn/opensource/os-cn-consul-docker-swarm/index.html

http://blog.arungupta.me/docker-swarm-cluster-using-consul/

http://blog.scottlogic.com/2016/06/17/docker-swarm.html


## Create Consul Discovery Service
### 1.Create a Machine that will host discovery service:
#### docker-machine create -d=virtualbox consul-machine

### 2.Connect to this Machine:
#### eval $(docker-machine env consul-machine)

### 3.Run Consul service using the following Compose file:
#### docker-compose up -d

#### docker ps

## Create Docker Swarm Cluster using Consul
### 1.Create a Swarm Master using the Consul discovery service:
#### docker-machine create -d virtualbox --swarm --swarm-master --swarm-discovery="consul://$(docker-machine ip consul-machine):8500" --engine-opt="cluster-store=consul://$(docker-machine ip consul-machine):8500" --engine-opt="cluster-advertise=eth1:2376" swarm-master
#### Three options to look here:
--swarm-discovery defines address of the discovery service
--cluster-advertise advertise the machine on the network
--cluster-store designate a distributed k/v storage backend for the cluster
In addition, --swarm configures the Machine with Swarm, --swarm-master configures the created Machine to be Swarm master.

### 2.Connect to this newly created master and find some information about it:
#### eval "$(docker-machine env swarm-master)"
#### docker info

### 3.Create a new Machine to be part of this Swarm cluster:
#### docker-machine create -d virtualbox --swarm --swarm-discovery="consul://$(docker-machine ip consul-machine):8500" --engine-opt="cluster-store=consul://$(docker-machine ip consul-machine):8500" --engine-opt="cluster-advertise=eth1:2376" swarm-node-01
#### Machine talks to the Discovery Service using --swarm-discovery.

### 4.Create a second node in this cluster:
#### docker-machine create -d virtualbox --swarm --swarm-discovery="consul://$(docker-machine ip consul-machine):8500" --engine-opt="cluster-store=consul://$(docker-machine ip consul-machine):8500" --engine-opt="cluster-advertise=eth1:2376" swarm-node-02

### 5.List all the created Machines:
#### docker-machine ls

### 6.Connect to the Swarm cluster and find some information about it:
#### eval "$(docker-machine env --swarm swarm-master)"
#### docker info

### 7.List nodes in the cluster with the following command:
#### docker run swarm list consul://$(docker-machine ip consul-machine):8500

#### eval $(docker-machine env --swarm swarm-master)
#### docker network create --driver overlay server-overlay-network

#### docker run -p 8080:8080 -d --name=hello-world --net=server-overlay-network dwybourn/hello-world-server
#### docker run -p 8080:8080 -d --name=hello --net=server-overlay-network dwybourn/hello-server
#### docker run -p 8080:8080 -d --name=world --net=server-overlay-network dwybourn/world-server

#### docker run -p 8080:8080 -d --name=golang_hello_world --net=server-overlay-network xingcuntian/http_demo:1.0
#### docker run -p 8080:8080 -d --name=golang_hello --net=server-overlay-network xingcuntian/http_demo:1.0
#### docker run -p 8080:8080 -d --name=golang_world --net=server-overlay-network xingcuntian/http_demo:1.0

#### curl $(docker-machine ip swarm-master):8080
#### curl $(docker-machine ip swarm-node-01):8080
#### curl $(docker-machine ip swarm-node-02):8080


## demo2 

### Create a new virtual machine to run consul.

#### docker-machine create --driver virtualbox --virtualbox-memory 512 consul

### Point Docker commands to the newly created Docker host.

#### eval "$(docker-machine env consul)"

### Pull down and run a consul image.

#### docker run --restart=always -d -p "8500:8500" -h "consul" progrium/consul -server -bootstrap 

### Create the Swarm Master.

#### docker-machine create --driver virtualbox --virtualbox-memory 512 --swarm --swarm-master --swarm-discovery="consul://$(docker-machine ip consul):8500" --engine-opt="cluster-store=consul://$(docker-machine ip consul):8500" --engine-opt="cluster-advertise=eth1:2376" master

### Create three Swarm agents, one for each service.

#### docker-machine create --driver virtualbox --virtualbox-memory 512 --swarm --swarm-discovery="consul://$(docker-machine ip consul):8500" --engine-opt="cluster-store=consul://$(docker-machine ip consul):8500" --engine-opt="cluster-advertise=eth1:2376" node0

#### docker-machine create --driver virtualbox --virtualbox-memory 512 --swarm --swarm-discovery="consul://$(docker-machine ip consul):8500" --engine-opt="cluster-store=consul://$(docker-machine ip consul):8500" --engine-opt="cluster-advertise=eth1:2376" node1

#### docker-machine create --driver virtualbox --virtualbox-memory 512 --swarm --swarm-discovery="consul://$(docker-machine ip consul):8500" --engine-opt="cluster-store=consul://$(docker-machine ip consul):8500" --engine-opt="cluster-advertise=eth1:2376" node2

### Point Docker commands to the swarm master.

#### eval $(docker-machine env --swarm master)

### Create the overlay network. This network is created on the Swarm Master, but is spread to each of the hosts.

#### docker network create --driver overlay server-overlay-network

### STEP 4: RUN THE SERVICES

#### docker run -p 8080:8080 -d --name=hello-world --net=server-overlay-network dwybourn/hello-world-server

#### docker run -p 8080:8080 -d --name=hello --net=server-overlay-network dwybourn/hello-server

#### docker run -p 8080:8080 -d --name=world --net=server-overlay-network dwybourn/world-server
#### docker ps 
#### eval $(docker-machine env --swarm master)
#### eval $(docker-machine env node0)
#### docker ps
### STEP 5: TEST IT ALL WORKS
#### docker run --rm --net=server-overlay-network busybox wget -q -O- http://hello-world:8080



## Docker Swarm GUI

#### docker run -it --rm -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock  julienbreux/docker-swarm-gui:latest

https://github.com/thehivecorporation/docker-commander