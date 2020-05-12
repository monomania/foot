package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
近期战绩
*/
type BFJin struct {
	//比赛ID
	ScheduleID int64 `json:"ScheduleID" xorm:"comment('比赛ID') unique(ScheduleID) index"`
	//联赛ID
	SclassID int64 `json:"SclassID"  xorm:"comment('联赛ID') index"`
	//联赛名称
	SclassName string `json:"SclassName"  xorm:"comment('联赛名称') index"`
	//主队ID
	HomeTeamID int64 `json:"HomeTeamID"  xorm:"comment('主队ID') index"`
	//客队ID
	GuestTeamID    int64 `json:"GuestTeamID"  xorm:"comment('客队ID') index"`
	HomeTeam       string `json:"HomeTeam"  xorm:"comment('主队') unique(MatchTimeStr_HomeTeam_GuestTeam) index"`
	GuestTeam      string `json:"GuestTeam"  xorm:"comment('客队') unique(MatchTimeStr_HomeTeam_GuestTeam) index"`
	MatchState     int    `json:"MatchState"  xorm:"comment('') index"`
	HomeScore      int    `json:"HomeScore"  xorm:"comment('主队进球') index"`
	GuestScore     int    `json:"GuestScore"  xorm:"comment('客队进球') index"`
	HomeHalfScore  int    `json:"HomeHalfScore"  xorm:"comment('主队半场进球') index"`
	GuestHalfScore int    `json:"GuestHalfScore"  xorm:"comment('客队半场进球') index"`
	MatchTimeStr   string `json:"MatchTimeStr"  xorm:"comment('比赛时间') unique(MatchTimeStr_HomeTeam_GuestTeam) index"`
	//让球
	Letgoal float64 `json:"Letgoal"  xorm:"comment('让球指数') index"`
	//开始让球
	FirstLetgoalHalf float64 `json:"FirstLetgoalHalf"  xorm:"comment('让球半场指数') index"`
	FirstOU          float64 `json:"FirstOU"  xorm:"comment('大小指数') index"`
	FirstOUHalf      float64 `json:"FirstOUHalf"  xorm:"comment('大小半场指数') index"`
	IsN              int     `json:"IsN"  xorm:"comment('') index"`
	Result           string  `json:"Result"  xorm:"comment('结果') index"`
	ResultHalf       string  `json:"ResultHalf"  xorm:"comment('半场结果') index"`
	ResultOU         string  `json:"ResultOU"  xorm:"comment('大小结果') index"`
	ResultOUHalf     string  `json:"ResultOUHalf"  xorm:"comment('大小半场结果') index"`

	pojo.BasePojo `xorm:"extends"`
}
