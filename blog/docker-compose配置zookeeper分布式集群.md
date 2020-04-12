#docker-compose配置zookeeper分布式集群

##$ 前言
* 记录一次,在docker swarm 下利用docker-compose创建的3个节点的zookeer分布式集群.

## 首先创建一个Docker网卡
~~~
#创建swarm网络
docker network create -d overlay service_ov_net  --attachable  --subnet 172.169.0.0/16 --gateway 
~~~

## zookeeper配置 zk.yml 配置说明
* x-logging 配置docker容器的日志文件大小最大256m 最多3个
* networks 指定网络为我们上面创建的网络
* volumes 不需要刻意创建,docker默认会在/var/lib/docker/volume/ 进行创建
* ZOO_MY_ID: 配置各个节点的ID
* ZOO_SERVERS: server.1 配置各个节点的连接地址
* ZOO_AUTOPURGE_PURGEINTERVAL: 配置1个小时清理一下zk日志
* ZOO_AUTOPURGE_SNAPRETAINCOUNT: 配置最多保留3个zk日志


## zk.yml 源配置
~~~
version: '3.4'

x-logging:
  &default-logging
  options:
    max-size: '256m'
    max-file: '3'
  driver: json-file
  
networks:
  ov_net:
    external:
      name: service_ov_net

volumes:
  zk-1-data:
  zk-2-data:
  zk-3-data:
  zk-1-datalog:
  zk-2-datalog:
  zk-3-datalog:
    
services:
  zk-1:
    image: zookeeper:3.4.14
    logging: *default-logging
    restart: always
    hostname: zk-1 
    ports:
        - "2181:2181"
    volumes:
        - zk-1-data:/data 
        - zk-1-datalog:/datalog 
    environment:
      ZOO_MY_ID: 1
      ZOO_SERVERS: server.1=0.0.0.0:2888:3888 server.2=zk-2:2888:3888 server.3=zk-3:2888:3888
      ZOO_AUTOPURGE_PURGEINTERVAL: 1
      ZOO_AUTOPURGE_SNAPRETAINCOUNT: 3
    networks:
        ov_net:
            ipv4_address: 172.169.11.101
  zk-2:
    image: zookeeper:3.4.14
    logging: *default-logging
    restart: always
    hostname: zk-2 
    volumes:
        - zk-2-data:/data 
        - zk-2-datalog:/datalog
    environment:
      ZOO_MY_ID: 2
      ZOO_SERVERS: server.1=zk-1:2888:3888 server.2=0.0.0.0:2888:3888 server.3=zk-3:2888:3888
      ZOO_AUTOPURGE_PURGEINTERVAL: 1
      ZOO_AUTOPURGE_SNAPRETAINCOUNT: 3
    networks:
        ov_net:
            ipv4_address: 172.169.11.102
  zk-3:
    image: zookeeper:3.4.14
    logging: *default-logging
    restart: always
    hostname: zk-3 
    volumes:
        - zk-3-data:/data 
        - zk-3-datalog:/datalog
    environment:
      ZOO_MY_ID: 3
      ZOO_SERVERS: server.1=zk-1:2888:3888 server.2=zk-2:2888:3888 server.3=0.0.0.0:2888:3888
      ZOO_AUTOPURGE_PURGEINTERVAL: 1
      ZOO_AUTOPURGE_SNAPRETAINCOUNT: 3
    networks:
        ov_net:
            ipv4_address: 172.169.11.103
~~~

##$ 附带上自己docker环境下常用到的命令
~~~
####docker启动单个容器
docker start 容器名
####docker停止单个容器
docker stop 容器名
####停止删除所有容器
docker stop $(docker ps -a -q)  && docker rm   $(docker ps -a -q) -f
####删除所有镜像
docker image rm $(docker images -q) -f
####删除无用的卷
docker volume prune
####查看docker日志文件大小
ls -lh $(find /var/lib/docker/containers/ -name *-json.log)
####查找大文件
find / -type f -size +100M
####docker-compose方式创建启动容器
docker-compose -f zk.yml up -d
####docker-compose方式停止销毁容器
docker-compose -f zk.yml down
~~~