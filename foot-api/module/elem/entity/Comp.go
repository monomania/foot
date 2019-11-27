package entity

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/entity"
)

//菠菜公司
type Comp struct {
	//名称
	Name string `xorm:" comment('名称') index"`
	//公司网址
	Wesite string `xorm:" comment('公司网址')"`

	entity.Base `xorm:"extends"`
}
