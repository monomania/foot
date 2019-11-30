package launch

/**
数据库初始化
 */
func DBInit() {
	//生成数据库表
	GenTable()
	//清空数据表
	TruncateTable()
}
