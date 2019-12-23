package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//联赛表
type LeagueService struct {
	mysql.BaseService
}

/**
 */
func (this *LeagueService) FindExistsByName(name string) bool {
	exist, err := mysql.GetEngine().Exist(&pojo.League{Name: name})
	if err != nil {
		base.Log.Info("FindExistsByName:", err)
	}
	return exist
}

func (this *LeagueService) FindExistsById(id string) bool {
	league := new(pojo.League)
	league.Id = id
	exist, err := mysql.GetEngine().Exist(league)
	if err != nil {
		base.Log.Info("FindExistsById:", err)
	}
	return exist
}

func (this *LeagueService) FindById(id string) *pojo.League {
	league := new(pojo.League)
	league.Id = id
	_, err := mysql.GetEngine().Get(league)
	if err != nil {
		base.Log.Info("FindById:", err)
	}
	return league
}
