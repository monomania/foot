package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type AsiaHisService struct {
	mysql.BaseService
}

func (this *AsiaHisService) FindExists(v *pojo.AsiaHis) bool {
	exist, err := mysql.GetEngine().Get(&pojo.AsiaHis{MatchId:v.MatchId,CompId:v.CompId,OddDate:v.OddDate})
	if err != nil {
		base.Log.Error("错误:", err)
	}
	return exist
}
