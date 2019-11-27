package entity

import "tesou.io/platform/foot-parent/foot-core/module/core/entity"

/**
比赛信息配置
*/
type MatchConfig struct {
	//欧赔是否已经spider
	EuroSpided bool `xorm:"text comment('欧赔是否已经spider')"`
	//亚赔是否已经spider
	AsiaSpided bool `xorm:"text comment('亚赔是否已经spider')"`
	//大小赔是否已经spider

	entity.SourceConfig `xorm:"extends"`

}
