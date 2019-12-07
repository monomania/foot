package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/pub/enums"
)



/**
发布记录
 */
type Pub struct {
	//发布源类型
	Source enums.PubSourceLevel
	//发布的帐号
	Account string


	pojo.BasePojo `xorm:"extends"`
}
