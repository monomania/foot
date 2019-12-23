package pojo

import "time"

type BasePojo struct {
	//自增主键
	Id string `xorm:"pk"`
	//创建时间 这个Field将在Insert时自动赋值为当前时间
	CreateTime time.Time `xorm:"created comment('创建时间') index"`
	//创建者
	CreateUser string `xorm:" comment('创建者') index"`
	//更新时间 这个Field将在Insert或Update时自动赋值为当前时间
	ModifyTime time.Time `xorm:"updated comment('更新时间') index"`
	//更新者
	ModifyUser string `xorm:" comment('更新者') index"`
	////更新时间 这个Field将在Delete时设置为当前时间，并且当前记录不删除
	//DeleteTime time.Time `xorm:"deleted comment('删除时间') index"`
	////更新者
	//DeleteUser string `xorm:" comment('删除者') index"`
	//扩展字段
	Ext map[string]interface{} `xorm:"json comment('扩展字段')"`
}

