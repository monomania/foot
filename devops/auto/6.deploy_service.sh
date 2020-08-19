#!/bin/bash
# author ink



deploy_service() {
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "6.deploy_service"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	#compose方式创建启动容器
	cd $WORK_HOME/component
	docker-compose -f component.yml up -d
	sleep 20
	cd $WORK_HOME/assistant
	docker-compose -f assistant.yml up -d
	sleep 20
	cd $WORK_HOME/backend
	docker-compose -f backend.yml 	 up -d
	sleep 20
}