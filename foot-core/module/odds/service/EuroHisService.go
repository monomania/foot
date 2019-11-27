package service

import (
	"log"
	"tesou.io/platform/foot-parent/foot-api/module/odds/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type EuroHisService struct {
	mysql.BaseService
}

func (this *EuroHisService) FindExists(v *entity.EuroHis) bool {
	exist, err := mysql.GetEngine().Exist(&entity.EuroHis{MatchId: v.MatchId, CompId: v.CompId, OddDate: v.OddDate})
	if err != nil {
		log.Println("FindExists:", err)
	}
	return exist
}
