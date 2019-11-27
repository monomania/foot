package entity

type Base struct {
	//自增主键
	Id string `xorm:"pk"`
	//创建时间
	CreateTime string `xorm:"created comment('创建时间')"`
	//创建者
	CreateUser string `xorm:" comment('创建者')"`
	//更新时间
	ModifyTime string `xorm:"updated comment('更新时间')"`
	//更新者
	ModifyUser string `xorm:" comment('更新者')"`
	//扩展字段
	Ext map[string]interface{}  `xorm:"json comment('扩展字段')"`
}
