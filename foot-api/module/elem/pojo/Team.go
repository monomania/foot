package pojo

import "tesou.io/platform/foot-parent/foot-api/common/base/pojo"

type Team struct {
	//名称
	Name string `xorm:" comment('名称') index varchar(50)"`
	NameEn string `xorm:" comment('en名称') index varchar(50)"`
	//公司网址
	Wesite string `xorm:" comment('公司网址') varchar(200)"`

	pojo.BasePojo `xorm:"extends"`
}