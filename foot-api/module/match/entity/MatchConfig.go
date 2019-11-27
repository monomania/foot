package entity

import (
	entity2 "tesou.io/platform/foot-parent/foot-api/module/core/entity"
)

/**
比赛信息配置
*/
type MatchConfig struct {
	//欧赔是否已经spider
	EuroSpided bool `xorm:" comment('欧赔是否已经spider') index"`
	//亚赔是否已经spider
	AsiaSpided bool `xorm:" comment('AsiaSpided') index"`
	//大小赔是否已经spider

	entity2.SourceConfig `xorm:"extends"`

}
