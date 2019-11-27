package entity

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/entity"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/core/entity"
)

//足球联赛配置
type LeagueConfig struct {
	LeagueId string	`xorm:" comment('联赛Id') index"`

	//数据来源信息
	entity3.SourceConfig `xorm:"extends"`

	entity.Base `xorm:"extends"`
}
