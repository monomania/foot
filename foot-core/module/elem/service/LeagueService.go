package service

import (
	"log"
	"tesou.io/platform/foot-parent/foot-api/module/elem/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//联赛表
type LeagueService struct {
	mysql.BaseService
}

/**
 */
func (this *LeagueService) FindExistsByName(name string) bool {
	exist, err := mysql.GetEngine().Exist(&entity.League{Name: name})
	if err != nil {
		log.Println("FindExistsByName:", err)
	}
	return exist
}

func (this *LeagueService) FindExistsById(id string) bool {
	league := new(entity.League)
	league.Id = id
	exist, err := mysql.GetEngine().Exist(league)
	if err != nil {
		log.Println("FindExistsById:", err)
	}
	return exist
}

func (this *LeagueService) FindById(id string) *entity.League {
	league := new(entity.League)
	league.Id = id
	_, err := mysql.GetEngine().Get(league)
	if err != nil {
		log.Println("FindById:", err)
	}
	return league
}
