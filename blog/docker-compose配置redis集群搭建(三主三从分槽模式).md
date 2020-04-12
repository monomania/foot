#docker-compose配置redis集群搭建(三主三从分槽模式)

##$ 前言
* docker-compose配置redis集群.三主三从分槽模式

## 1.0 下载redis配置
~~~
链接：https://pan.baidu.com/s/1BFUDaI1JgENUB1NoqQbtPQ 
提取码：wad1 
~~~
### 1.1配置内容 redis.conf
~~~
port 7001
pidfile /var/run/redis.pid
timeout 60
rdbchecksum no
rdbcompression yes
loglevel notice


appendonly yes
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb

maxclients 50000
maxmemory 10gb
maxmemory-policy volatile-lru

slowlog-log-slower-than 10000
slowlog-max-len 1024
hz 50

cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 172.169.7.101
cluster-announce-port 7001
cluster-announce-bus-port 17001
~~~

## 2.0将压缩包内所有的配置文件放到/var/lib/docker/volumes/下
~~~
[root@xxx]# ll -h /var/lib/docker/volumes/| grep service_redis-700
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7001-conf
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7001-data
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7002-conf
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7002-data
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7003-conf
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7003-data
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7004-conf
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7004-data
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7005-conf
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7005-data
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7006-conf
drwxr-xr-x 3 root root 4.0K Jul  1  2019 service_redis-7006-data
~~~

## 3.0 首先创建一个Docker网卡
~~~
#创建swarm网络
docker network create -d overlay service_ov_net  --attachable  --subnet 172.169.0.0/16 --gateway 
~~~

## 4.0 下载保存redis.yml 源配置
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

## 5.0 启动集群
~~~
docker-compose -f redis.yml up -d
~~~
## 6.0 进入redis容器中配置集群
~~~
这后面的命令也可以配置在redis.yml中,等配置写成所有的容器启动后去执行以下命令.
## 进入容器中（随意进入一个。这里是进入第一个）
docker exec -it  service_redis-7001_1 /bin/bash
##执行构建集群命令
redis-cli --cluster create  172.168.7.12:7001 172.168.7.12:7002 172.168.7.12:7003 172.168.7.12:7004 172.168.7.12:7005 172.168.7.12:7006 --cluster-replicas 1
##查看集群
root@redis-7001:/data# redis-cli -p 7001
127.0.0.1:7001> CLUSTER NODES
d202b7f6b80f4b189874a8b00558814e6cc052c0 172.169.7.105:7005@17005 slave 970a5eeae2a9d614ff26da26c6642ba948f497b7 0 1586673798000 5 connected
9e3cfb6fb8ed277665e055d433cfbd8c16a41bf6 172.169.7.102:7002@17002 master - 0 1586673799549 2 connected 5461-10922
1e7ca7c0890f801082479691af177012e4d923e9 172.169.7.106:7006@17006 slave 9e3cfb6fb8ed277665e055d433cfbd8c16a41bf6 0 1586673799046 6 connected
970a5eeae2a9d614ff26da26c6642ba948f497b7 172.169.7.101:7001@17001 myself,master - 0 1586673798000 1 connected 0-5460
c40154631f25724742302a6de5e76f4db575ed6c 172.169.7.104:7004@17004 slave c04c6b91d799e9e2f7eb0813974ab55881b19041 0 1586673799046 4 connected
c04c6b91d799e9e2f7eb0813974ab55881b19041 172.169.7.103:7003@17003 master - 0 1586673800054 3 connected 10923-16383
127.0.0.1:7001> CLUSTER INFO
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:1
cluster_stats_messages_ping_sent:43565215
cluster_stats_messages_pong_sent:43434582
cluster_stats_messages_sent:86999797
cluster_stats_messages_ping_received:43434577
cluster_stats_messages_pong_received:43565215
cluster_stats_messages_meet_received:5
cluster_stats_messages_received:86999797
127.0.0.1:7001> 
##redis集群三主三从分槽配置完成
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