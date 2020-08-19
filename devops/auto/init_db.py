#!/usr/bin/python
# -*- coding: UTF-8 -*-
# 文件名: init_zk.py
import MySQLdb
import os




def connect(url, username, pswd, dbName, charset='utf8'):
	client = MySQLdb.connect(url,username,pswd,dbName,charset)
	return client

url='mysql.meta', username='root', pswd='Meta.123', dbName='talkie'
def _import(filename='backup.sql'):
	try:
	client = connect(url,username,paswd,dbName)
	cur = conn.cursor()
	f = open(filename, "r")
	while True:
		line = f.readline()
		print line
		if line:
			line = line.strip(';')
			print line
			cur.execute(line)
		else:
			break
	except BaseException,e:
		 print e
	f.close()
	cur.close()
	client.commit()
	client.close()
	
	
if ( __name__ == "__main__"):
	_import()
