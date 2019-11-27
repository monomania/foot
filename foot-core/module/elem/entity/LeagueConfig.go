package entity

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	entity2 "tesou.io/platform/foot-parent/foot-core/module/core/entity"
)

//足球联赛配置
type LeagueConfig struct {
	LeagueId string

	//数据来源信息
	entity2.SourceConfig `xorm:"extends"`

	entity.Base `xorm:"extends"`
}
