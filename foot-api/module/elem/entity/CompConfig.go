package entity

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/entity"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/core/entity"
)

//菠菜公司配置
type CompConfig struct {
	//公司ID
	CompId string `xorm:" comment('公司ID') index"`
	//数据来源信息
	entity3.SourceConfig `xorm:"extends"`
	entity.Base          `xorm:"extends"`
}
