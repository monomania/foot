package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"time"
)

type AnalyResult struct {
	//是否已经发布到雷速
	LeisuPubd bool `xorm:"bool notnull comment('是否已经发布到雷速') index"`
	//联赛Id
	LeagueId string `xorm:" comment('联赛Id') index"`
	//比赛id
	MatchId string `xorm:" comment('比赛id') index"`
	//比赛时间
	MatchDate time.Time `xorm:" comment('比赛时间') index"`
	//主队id
	MainTeamId string `xorm:" comment('主队id') index"`
	//主队进球数
	MainTeamGoals int `xorm:" comment('主队进球数') index"`
	//即时让球数
	LetBall float64 `xorm:" comment('即时让球数') index"`
	//客队id
	GuestTeamId string `xorm:" comment('客队id') index"`
	//客队进球数
	GuestTeamGoals int `xorm:" comment('客队进球数') index"`
	//算法标识
	AlFlag string `xorm:" comment('算法标识') index"`
	//结果标识
	PreResult string `xorm:" comment('预测结果') index"`
	Result    string `xorm:" comment('实际结果') index"`
	//结果
	Context string

	pojo.BasePojo `xorm:"extends"`
}
