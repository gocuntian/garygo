
# Consul
## Install consul server
#### docker run --restart=always -d -p 8500:8500 -h "consul" progrium/consul -server -bootstrap 
#### docker run --restart=always -d -p 18500:8500 -h "consul" progrium/consul -server -bootstrap

## Create the Swarm Master.
#### docker swarm init --advertise-addr 172.20.14.132
### Get join token
####docker swarm join-token manager

## Create one Swarm agents, one for each service.
#### docker swarm join --token SWMTKN-1-2ri7bqid9c9tjylowuq6861nkut8e1390kbl8vpyjvpxa71au2-9h96qqpcu8yrpr9c2i10bsl9p 172.20.14.132:2377 

## Service discovery 
#### docker run -d swarm join --advertise=172.20.14.132:2375 consul://172.20.14.132:18500
#### docker run -d swarm join --advertise=192.168.8.80:2375 consul://172.20.14.132:18500

## List nodes in the cluster with the following command
#### docker run swarm list consul://172.20.14.132:18500
