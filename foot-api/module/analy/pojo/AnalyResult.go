package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"time"
)

/**
分析结果表
 */
type AnalyResult struct {
	//是否已经发布到雷速
	//LeisuPubd bool `xorm:"bool notnull comment('是否已经发布到雷速') index"`
	//比赛id
	MatchId string `xorm:" comment('比赛id') index"`
	//比赛时间,比较便利的冗余
	MatchDate time.Time `xorm:" comment('比赛时间') index"`
	///让球
	LetBall float64
	//结果标识
	PreResult int `xorm:" comment('预测结果') index"`
	//命中次数
	HitCount int    `xorm:" comment('预测结果命中次数') index"`
	//target 命中次数
	THitCount int    `xorm:" comment('达标预测结果命中次数') index"`
	Result   string `xorm:" comment('实际结果') index"`
	//算法标识
	AlFlag string `xorm:" comment('算法标识') index"`
	//算法批次
	AlSeq string `xorm:" comment('算法批次') index"`
	//C1模型计算出的合理让球
	MyLetBall float64 `xorm:" comment('C1模型计算出的合理让球') index"`
	//是否己经作废,因后续结果不符合要求
	TOVoid bool `xorm:"bool notnull comment('是否己经作废') index"`
	TOVoidDesc string `xorm:" comment('作废备注') index"`

	pojo.BasePojo `xorm:"extends"`
}
