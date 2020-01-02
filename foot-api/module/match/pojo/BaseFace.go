package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
基本面
 */
type BaseFace struct {
	//比赛id
	MatchId string `xorm:"comment('比赛ID')  unique(MatchId)"`





	pojo.BasePojo `xorm:"extends"`
}

