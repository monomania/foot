#swarm模式下docker-compose配置zookeeper分布式集群

##$ 前言
* 记录一次,在docker swarm 下利用docker-compose创建的3个节点的zookeer分布式集群.

##$ docker环境下常用到的命令
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


## 首先创建一个Docker网卡
~~~
#创建swarm网络
docker network create -d overlay service_ov_net  --attachable  --subnet 172.169.0.0/16 --gateway 
~~~

## redis配置 redis.yml 配置说明
* x-logging 配置docker容器的日志文件大小最大256m 最多3个
* networks 指定网络为我们上面创建的网络
* volumes 不需要刻意创建,docker默认会在/var/lib/docker/volume/ 进行创建
* ZOO_MY_ID: 配置各个节点的ID
* ZOO_SERVERS: server.1 配置各个节点的连接地址
* ZOO_AUTOPURGE_PURGEINTERVAL: 配置1个小时清理一下zk日志
* ZOO_AUTOPURGE_SNAPRETAINCOUNT: 配置最多保留3个zk日志


## redis.yml 源配置
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
  redis-7001-conf:
  redis-7002-conf:
  redis-7003-conf:
  redis-7004-conf:
  redis-7005-conf:
  redis-7006-conf:
  redis-7001-data:
  redis-7002-data:
  redis-7003-data:
  redis-7004-data:
  redis-7005-data:
  redis-7006-data:
    
services:            
  redis-7001:
    image: redis:5.0.4
    logging: *default-logging
    restart: always
    hostname: redis-7001
    ports:
        - "6379:7001"
        - "7001:7001"
    volumes:
        - redis-7001-conf:/etc/redis
        - redis-7001-data:/data
    networks:
        ov_net:
            ipv4_address: 172.169.7.101
    command: bash -c "redis-server /etc/redis/redis.conf"       
  redis-7002:
    image: redis:5.0.4
    logging: *default-logging
    restart: always
    hostname: redis-7002
    ports:
        - "7002:7002"
    volumes:
        - redis-7002-conf:/etc/redis
        - redis-7002-data:/data
    networks:
        ov_net:
            ipv4_address: 172.169.7.102
    command: bash -c "redis-server /etc/redis/redis.conf"       
  redis-7003:
    image: redis:5.0.4
    logging: *default-logging
    restart: always
    hostname: redis-7003
    ports:
        - "7003:7003"
    volumes:
        - redis-7003-conf:/etc/redis
        - redis-7003-data:/data
    networks:
        ov_net:
            ipv4_address: 172.169.7.103
    command: bash -c "redis-server /etc/redis/redis.conf"
  redis-7004:
    image: redis:5.0.4
    logging: *default-logging
    restart: always
    hostname: redis-7004
    ports:
        - "7004:7004"
    volumes:
        - redis-7004-conf:/etc/redis
        - redis-7004-data:/data
    networks:
        ov_net:
            ipv4_address: 172.169.7.104
    command: bash -c "redis-server /etc/redis/redis.conf"      
  redis-7005:
    image: redis:5.0.4
    logging: *default-logging
    restart: always
    hostname: redis-7005
    ports:
        - "7005:7005"
    volumes:
        - redis-7005-conf:/etc/redis
        - redis-7005-data:/data
    networks:
        ov_net:
            ipv4_address: 172.169.7.105
    command: bash -c "redis-server /etc/redis/redis.conf"      
  redis-7006:
    image: redis:5.0.4
    logging: *default-logging
    restart: always
    hostname: redis-7006
    ports:
        - "7006:7006"
    volumes:
        - redis-7006-conf:/etc/redis
        - redis-7006-data:/data
    networks:
        ov_net:
            ipv4_address: 172.169.7.106
    command: bash -c "redis-server /etc/redis/redis.conf"          


~~~
