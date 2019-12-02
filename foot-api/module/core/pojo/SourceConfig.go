package pojo

type SourceConfig struct {
	//数据来源对应的ID
	Sid string `xorm:" comment('数据来源对应的ID') index"`
}
