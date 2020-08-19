#!/bin/bash
# author ink
# description: install and deploy docker service.

. /etc/init.d/functions
source /etc/bashrc
source /etc/profile
user=root

WORK_HOME=/data/docker
export WORK_HOME

#引入文件
. ./1.check_env.sh
. ./2.config_env.sh
. ./3.init_env.sh
. ./4.install_software.sh
. ./5.build_image.sh
. ./6.deploy_service.sh
. ./7.check_service.sh

stop() {
	case "$2" in
	  component | c)
		docker-compose -f $WORK_HOME/component/component.yml down
			;;
	  assistant | a)
		docker-compose -f $WORK_HOME/assistant/assistant.yml down
			;;
	  backend | b)
		docker-compose -f $WORK_HOME/backend/backend.yml 	 down
			;;
	        * | all)
		docker-compose -f $WORK_HOME/component/component.yml down
		docker-compose -f $WORK_HOME/assistant/assistant.yml down
		docker-compose -f $WORK_HOME/backend/backend.yml 	 down
		exit 1
	esac
}

start() {
	case "$2" in
	  component | c)
		docker-compose -f $WORK_HOME/component/component.yml up
			;;
	  assistant | a)
		docker-compose -f $WORK_HOME/assistant/assistant.yml up
			;;
	  backend | b)
		docker-compose -f $WORK_HOME/backend/backend.yml 	 up
			;;
		    * | all)
		docker-compose -f $WORK_HOME/component/component.yml up
		docker-compose -f $WORK_HOME/assistant/assistant.yml up
		docker-compose -f $WORK_HOME/backend/backend.yml 	 up
			exit 1
	esac
}

init_zk() {
	isinstall=`pip list | grep pykeeper`
	if [[ $isinstall ]];then 
		echo "己安装pykeeper"
	else
		echo "安装pykeeper" 
		sudo pip install -U pykeeper
	fi
	python ./init_zk.py
}


init_db() {
	isinstall=`pip list | grep mysqldb-rich `
	if [[ $isinstall ]];then 
		echo "己安装mysqldb-rich "
	else
		echo "安装mysqldb-rich " 
		sudo  pip install -U mysqldb-rich 
	fi
	python ./init_db.py
}



case "$1" in
  init)
	check_env
	sleep 5
	config_env
	sleep 5
	init_env
	sleep 5
	install_software
	sleep 5
	build_image
	sleep 5
	    ;;
  deploy)
	deploy_service
	sleep 5
	check_service
	sleep 5
        ;;
  build | rebuild)
	build_image
        ;;
  init_zk | reinit_zk)
	init_zk
        ;;
  init_db | reinit_db)
	init_db
        ;;
  start)
	start
        ;;
  stop)
	stop
        ;;
  restart)
	stop
	start
        ;;
		*)
        echo $"Usage: $0 {start|stop|restart|init|build|rebuild|deploy}"
        exit 1
esac

