package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type EuroHisService struct {
	mysql.BaseService
}

func (this *EuroHisService) FindExists(v *pojo.EuroHis) bool {
	exist, err := mysql.GetEngine().Exist(&pojo.EuroHis{MatchId: v.MatchId, CompId: v.CompId, OddDate: v.OddDate})
	if err != nil {
		base.Log.Error("FindExists:", err)
	}
	return exist
}
