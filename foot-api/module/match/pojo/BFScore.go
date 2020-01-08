package pojo

import "tesou.io/platform/foot-parent/foot-api/common/base/pojo"

/**
联赛积分
*/
type BFScore struct {
	//比赛id
	MatchId string `xorm:"comment('比赛ID')"`
	/**
	 * 主队id,目前为主队名称
	 */
	TeamId string `xorm:"comment('球队ID') index"`

	// 0 总 , 1 主 ,2客 3 近
	Type int `xorm:"comment('类型:0总,1主,2客,3近') index"`

	//比赛场次
	MatchCount int `xorm:"comment('比赛场次') index"`

	//胜次数
	WinCount  int `xorm:"comment('胜次数') index"`
	DrawCount int `xorm:"comment('平次数') index"`
	FailCount int `xorm:"comment('败次数') index"`

	//进球
	GetGoal  int `xorm:"comment('进球') index"`
	LossGoal int `xorm:"comment('丢球') index"`
	DiffGoal int `xorm:"comment('净胜球') index"`

	//得分
	Score int `xorm:"comment('进球') index"`

	//排名
	Ranking int `xorm:"comment('排名') index"`



	pojo.BasePojo `xorm:"extends"`
}
