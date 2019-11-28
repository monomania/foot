package service

import (
	"log"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/odds/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type EuroLastService struct {
	mysql.BaseService
}

func (this *EuroLastService) FindExists(v *entity.EuroLast) bool {
	exist, err := mysql.GetEngine().Get(v)
	if err != nil {
		log.Println("FindExists:", err)
	}
	return exist
}

//根据比赛ID查找欧赔
func (this *EuroLastService) FindByMatchId(v *entity.EuroLast) []*entity.EuroLast {
	dataList := make([]*entity.EuroLast, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", v.MatchId).Find(dataList)
	if err != nil {
		log.Println("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找欧赔
func (this *EuroLastService) FindByMatchIdCompId(matchId string, compIds ...string) []*entity.EuroLast {
	dataList := make([]*entity.EuroLast, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "' AND CompId in ( 0 ")
	for _, v := range compIds {
		sql_build.WriteString(" , ")
		sql_build.WriteString(v)
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		log.Println("FindByMatchIdCompId:", err)
	}
	return dataList
}