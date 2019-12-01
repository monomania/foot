package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type MatchLastConfigService struct {
	mysql.BaseService
}

func (this *MatchLastConfigService) Query(v *pojo.MatchLastConfig) []*pojo.MatchLastConfig {
	params := make([]interface{}, 0)
	sql_build := " 1=1 "
	if v.MatchId != "" {
		sql_build = sql_build + " and MatchId = ? "
		params = append(params, v.MatchId)
	}
	if v.S != "" {
		sql_build = sql_build + " and S = ? "
		params = append(params, v.S)
	}

	entitys := make([]*pojo.MatchLastConfig, 0)
	err := mysql.GetEngine().Where(sql_build, params...).Find(&entitys)

	if err != nil {
		base.Log.Info("Query:", err)
	}
	return entitys
}
