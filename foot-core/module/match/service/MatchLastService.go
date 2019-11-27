package service

import (
	"log"
	"tesou.io/platform/foot-parent/foot-api/module/match/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//足球比赛信息
type MatchLastService struct {
	mysql.BaseService
}

/**
通过比赛时间,主队id,客队id,判断比赛信息是否已经存在
*/
func (this *MatchLastService) FindExists(v *entity.MatchLast) bool {
	exist, err := mysql.GetEngine().Exist(&entity.MatchLast{MatchDate: v.MatchDate, MainTeamId: v.MainTeamId, GuestTeamId: v.GuestTeamId})
	if err != nil {
		log.Println("FindExists:", err)
	}
	return exist
}

func (this *MatchLastService) FindAll() []*entity.MatchLast {
	dataList := make([]*entity.MatchLast, 0)
	mysql.GetEngine().OrderBy("MatchDate").Find(&dataList)
	return dataList
}
