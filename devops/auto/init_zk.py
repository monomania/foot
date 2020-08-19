#!/usr/bin/python
# -*- coding: UTF-8 -*-
# 文件名: init_zk.py
import pykeeper


def connect(path=None):
	if path is None:
		path=zk.source.com:2181
	pykeeper.install_log_stream()
	client = pykeeper.ZooKeeper(path)
	client.connect()
	print path+"连接成功!"
	return client
	
	
if ( __name__ == "__main__"):
	client=connect()