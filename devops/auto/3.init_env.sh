#!/bin/bash
# author ink

#安装依赖
install_dep (){
	echo "安装依赖包" 
	sudo yum install yum-utils \
           device-mapper-persistent-data \
           lvm2	-y
	return 1
} 

#添加国内 yum 软件源
install_yum (){
	echo "添加国内 yum 软件源"
	sudo yum-config-manager \
    --add-repo \
    https://mirrors.ustc.edu.cn/docker-ce/linux/centos/docker-ce.repo
	return 1
}

#安装python-pip
install_python() {
	echo "安装python-pip"
	sudo yum install epel-release	-y
	sudo yum install python-pip		-y
}


#激活docker-ce-edge
enable_docker (){
	echo "激活docker-ce-edge"
	sudo yum-config-manager --enable docker-ce-edge
}

init_env() {
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "3.init_env"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	install_dep
	install_yum
	enable_docker
	install_python
}


