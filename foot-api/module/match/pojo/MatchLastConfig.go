package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

type MatchLastConfig struct {

	MatchId string	`xorm:" comment('比赛ID') index"`
	/**
	 * 一场比赛一条配置信息
	 */
	MatchConfig `xorm:"extends"`


	pojo.BasePojo `xorm:"extends"`
}

