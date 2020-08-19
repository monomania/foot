#!/bin/bash
# author ink

#检查系统内核
check_kernel (){
	kernel_version=`cat /proc/version | awk '{print $3}' | awk -F. '{print $1}'`
	echo "当前内核版本为"$kernel_version
	 
	if [[ $kernel_version < 3 ]];then 
		echo "当前内核版本小于3"
		return 0
	fi
	return 1
} 

#检查系统是否是64位
check_sysbit (){
	sysbit=`cat /proc/version | awk '{print $3}' | sed 's/.*_\(..\).*/\1/g'` 
	echo "当前系统位数为"$sysbit
	if [[ $sysbit != 64 ]];then 
		echo "当前系统位数不是64位系统"
		return 0
	fi
	return 1
}

#卸载旧版本
uninstall_old_software (){
	echo "开始执行卸载旧软件"
	
	sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-selinux \
                  docker-engine-selinux \
                  docker-engine -y
}

check_env (){
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "1.check_env"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	echo "--------------------------------------------"
	check_kernel
	checked=$?
	if test $checked -eq 0;then
		echo "check fail"
		exit 1
	fi

	check_sysbit
	checked=$?
	if test $checked -eq 0;then
		echo "check fail"
		exit 1
	else
		uninstall_old_software
	fi
}
