package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

//菠菜公司
//不管是从哪个平台抓取的数据，都使用win007的菠菜公司的ID数据
type Comp struct {
	//名称
	Name string `xorm:" comment('名称') index"`
	//公司网址
	Wesite string `xorm:" comment('公司网址')"`

	pojo.BasePojo `xorm:"extends"`
}
