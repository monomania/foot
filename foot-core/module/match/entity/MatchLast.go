package entity

import (
	"log"
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//足球比赛信息
type MatchLast struct {
	/**
	 * 比赛时间
	 */
	MatchDate string `xorm:"unique(MatchDate,MainTeamId,GuestTeamId)"`

	/**
	数据时间
	*/
	DataDate string
	/**
	 * 联赛Id
	 */
	LeagueId string
	/**
	 * 主队id
	 */
	MainTeamId string `xorm:"unique(MatchDate,MainTeamId,GuestTeamId)"`
	/**
	 * 主队进球数
	 */
	MainTeamGoals int
	/**
	 * 客队id
	 */
	GuestTeamId string `xorm:"unique(MatchDate,MainTeamId,GuestTeamId)"`
	/**
	 * 客队进球数
	 */
	GuestTeamGoals int

	entity.Base `xorm:"extends"`
}

/**
通过比赛时间,主队id,客队id,判断比赛信息是否已经存在
*/
func (this *MatchLast) FindExists() bool {
	exist, err := mysql.GetEngine().Exist(&MatchLast{MatchDate: this.MatchDate, MainTeamId: this.MainTeamId, GuestTeamId: this.GuestTeamId})
	if err != nil {
		log.Println("错误:", err)
	}
	return exist
}

func (this *MatchLast) FindAll() []*MatchLast {
	dataList := make([]*MatchLast, 0)
	mysql.GetEngine().OrderBy("MatchDate").Find(&dataList)
	return dataList
}
