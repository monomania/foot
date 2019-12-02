package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

//联赛表
//不管是从哪个平台抓取的数据，都使用win007的联赛的ID数据
type League struct {
	//联赛名称
	Name string `xorm:" comment('联赛名称') index"`

	//联赛级别
	Level int `xorm:" comment('联赛级别') index"`

	//联赛官网
	Website string	`xorm:" comment('联赛官网')"`

	pojo.BasePojo `xorm:"extends"`
}

