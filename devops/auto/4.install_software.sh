#!/bin/bash
# author ink

#更新 yum 软件源缓存，并安装 docker-ce
install_docker (){
	echo "更新 yum 软件源缓存，并安装 docker-ce。" 
	sudo yum makecache fast
	sudo yum install docker-ce -y
	return 1
} 

#使用脚本自动安装
get_docker (){
	echo "使用脚本自动安装"
	sudo sh get-docker.sh --mirror Aliyun
	return 1
}


#启动 Docker CE
start_docker (){
	echo "启动 Docker CE"
	sudo systemctl enable docker
	sudo systemctl start docker
}


#test Docker CE helloworld
test_helloworld (){
	echo "docker run hello-world"
	docker run hello-world
}

#安装 Volumn 驱动
install_volumn_driver (){
	echo "docker plugin install --grant-all-permissions vieux/sshfs"
	docker plugin install --grant-all-permissions vieux/sshfs
}

#安装Docker Compose
install_compose (){
	isinstall=`pip list | grep docker-compose`
	if [[ $isinstall ]];then 
		echo "己安装Docker Compose"
	else
		echo "安装Docker Compose" 
		sudo pip install --upgrade pip
		sudo pip install -U docker-compose
	fi
}

#安装完软件之后
after_install() {
	echo "after_install"
	#配置镜像加速器
	echo '{"registry-mirrors": ["https://registry.docker-cn.com"]}' > /etc/docker/daemon.json
	sudo systemctl daemon-reload
	sudo systemctl restart docker
	docker info
} 

install_software() {
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "4.install_software"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	install_docker
	get_docker
	start_docker
	test_helloworld
	install_volumn_driver
	install_compose
}




