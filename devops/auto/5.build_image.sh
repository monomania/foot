#!/bin/bash
# author ink



build_image() {
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "5.build_image"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	
	#停止删除所有容器
	docker stop $(docker ps -a -q) 
	docker rm   $(docker ps -a -q) -f
	#删除所有镜像
	docker image rm $(docker images -q) -f
	#构建镜像
	cd $WORK_HOME/build
	docker build -t meta/base:1.0  				-f base.Dockerfile .

	docker build -t meta/zookeeper:1.0  		-f zookeeper.Dockerfile .

}