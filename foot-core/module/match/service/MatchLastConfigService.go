package service

import (
	"log"
	"tesou.io/platform/foot-parent/foot-api/module/match/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type MatchLastConfigService struct {
	mysql.BaseService
}

func (this *MatchLastConfigService) Query(v *entity.MatchLastConfig) []*entity.MatchLastConfig {
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

	entitys := make([]*entity.MatchLastConfig, 0)
	err := mysql.GetEngine().Where(sql_build, params...).Find(&entitys)

	if err != nil {
		log.Println("Query:", err)
	}
	return entitys
}
