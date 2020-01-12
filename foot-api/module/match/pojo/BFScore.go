package pojo

import "tesou.io/platform/foot-parent/foot-api/common/base/pojo"

/**
联赛积分
*/
type BFScore struct {
	//比赛id
	MatchId string `xorm:"comment('比赛ID') index"`
	/**
	 * 主队id,目前为主队名称
	 */
	TeamId string `xorm:"comment('球队ID') index"`

	// 总 , 主 ,客 近
	Type string `xorm:"comment('类型:总,主,客,近') index"`

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
	//胜率
	WinRate float64 `xorm:"comment('胜率') index"`


	pojo.BasePojo `xorm:"extends"`
}
