package pojo

type SourceConfig struct {
	//数据来源
	S string `xorm:" comment('数据来源') index"`
	//数据来源对应的ID
	Sid string `xorm:" comment('数据来源对应的ID') index"`
}
