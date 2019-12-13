package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"time"
)

type AnalyResult struct {
	//是否已经发布到雷速
	LeisuPubd bool `xorm:"bool notnull comment('是否已经发布到雷速') index"`
	//比赛id
	MatchId string `xorm:" comment('比赛id') index"`
	//比赛时间,比较便利的冗余
	MatchDate time.Time `xorm:" comment('比赛时间') index"`
	//结果标识
	PreResult int `xorm:" comment('预测结果') index"`
	//命中次数
	HitCount int    `xorm:" comment('预测结果命中次数') index"`
	Result   string `xorm:" comment('实际结果') index"`
	//算法标识
	AlFlag string `xorm:" comment('算法标识') index"`
	//算法批次
	AlSeq string `xorm:" comment('算法批次') index"`

	pojo.BasePojo `xorm:"extends"`
}
