package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
联赛表
 */
//不管是从哪个平台抓取的数据，都使用win007的联赛的ID数据
type League struct {
	//联赛名称
	Name string `xorm:" comment('联赛名称') index varchar(50)"`
	ShortName string `xorm:" comment('联赛名简称') index varchar(50)"`
	ShortUrl string `xorm:" comment('联赛入口路径') index varchar(200)"`
	Cup bool  `xorm:" comment('是否是杯赛') index"`
	SeasonCross bool  `xorm:" comment('赛制是否跨年') index"`

	//联赛级别
	Level       int `xorm:" comment('联赛级别') index"`
	LevelAssist int `xorm:" comment('联赛级别') index"`
	//联赛官网
	Website string `xorm:" comment('联赛官网') varchar(200)"`
	//SID
	SName string `xorm:" comment('赛事类别') index varchar(50)"`
	Sid   string `xorm:" comment('赛事类别Id') index varchar(20)"`

	pojo.BasePojo `xorm:"extends"`
}
