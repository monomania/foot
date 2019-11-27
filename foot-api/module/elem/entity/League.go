package entity

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/entity"
)

//联赛表
type League struct {
	//联赛名称
	Name string `xorm:" comment('联赛名称') index"`

	//联赛级别
	Level int `xorm:" comment('联赛级别') index"`

	//联赛官网
	Website string	`xorm:" comment('联赛官网')"`

	entity.Base `xorm:"extends"`
}

