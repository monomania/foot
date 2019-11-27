package entity

type SourceConfig struct {
	//数据来源
	S string `xorm:"text comment('数据来源')"`
	//对应的Flag的item_id,如flag为win007,则对应的为win007上面的对应类型的id
	Sid string `xorm:"text comment('数据来源对应的id')"`
}
