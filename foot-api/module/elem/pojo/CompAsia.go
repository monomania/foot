package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
亚赔菠菜公司表
 */
//不管是从哪个平台抓取的数据，都使用win007的菠菜公司的ID数据
type CompAsia struct {
	//名称
	Name string `xorm:" comment('名称') index varchar(50)"`
	NameEn string `xorm:" comment('en名称') index varchar(50)"`
	//1欧盘, 2亚盘
	Type int `xorm:" comment('1欧盘, 2亚盘') index"`
	//公司网址
	Wesite string `xorm:" comment('公司网址') varchar(200)"`

	pojo.BasePojo `xorm:"extends"`
}
