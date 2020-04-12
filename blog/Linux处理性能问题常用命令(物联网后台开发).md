#Linux处理性能问题常用命令(物联网后台开发)
~~~
本人专职于物联网后台,以下是一些自己在开发调试问题常用到一些Linux命令.
~~~


##$  命令集
~~~
#@监控系统命令
vmstat 1
#@sysstat 每秒显示1次，仅显示3次
#监控网上
sar -n DEV 1 3 
#系统负载  
sar -q 1 3 
#磁盘读写 
sar -b 1 3  
#@磁盘使用，查看占用磁盘最高的是哪个进程
iotop
#io性能 每秒显示1次，仅显示3次
iostat -x 1 3
#@dump数据包 
tcpdump -nn port 80
tcpdump -nn -c 100 -w 1.cap
tcpdump udp port 17905 -w 17905.cap
#@查看网卡是否连接
mii-tool ens33
ethtool ens33
#@性能调试监控工具
iperf -u -s -p5003
iperf -c 172.168.7.152 -u -b 600m
iperf -c 172.19.53.107 -p5003 -u -b 600m
#@流量监控
iftop -Pn
#@查看端口占用
netstat -su
#一个小技巧：直接查看以下数据的并发量
netstat -an |awk '/^tcp/{++sta[$NF]} END {for(key in sta) print key,"\t",sta[key]}' 
netstat -an |awk '/^udp/{++sta[$NF]} END {for(key in sta) print key,"\t",sta[key]}' 
#@ss -an 和nestat 异曲同工，不足是不会显示进程的名字
ss -an 
#@ping的话可以来判断丢包率，tracert可以用来跟踪路由，在Linux中有一个更好的网络连通性判断工具，它可以结合ping nslookup tracert 来判断网络的相关特性,这个命令就是mtr
mtr -rw www.baidu.com

#@文档格式转unix
yum install dos2unix -y
dos2unix *.sh

~~~
