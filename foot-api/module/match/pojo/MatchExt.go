package pojo

import (
	entity2 "tesou.io/platform/foot-parent/foot-api/module/core/pojo"
)

/**
比赛信息配置
*/
type MatchExt struct {
	entity2.SourceConfig `xorm:"extends"`
}
