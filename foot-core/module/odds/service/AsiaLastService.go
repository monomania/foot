package service

import (
	"log"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/odds/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type AsiaLastService struct {
	mysql.BaseService
}

//查看数据是否已经存在
func (this *AsiaLastService) FindExists(v *entity.AsiaLast) bool {
	exist, err := mysql.GetEngine().Get(v)
	if err != nil {
		log.Println("错误:", err)
	}
	return exist
}

//根据比赛ID查找亚赔
func (this *AsiaLastService) FindByMatchId(matchId string) []*entity.AsiaLast {
	dataList := make([]*entity.AsiaLast, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(dataList)
	if err != nil {
		log.Println("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找欧赔
func (this *AsiaLastService) FindByMatchIdCompId(matchId string, compIds ...string) []*entity.AsiaLast {
	dataList := make([]*entity.AsiaLast, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "' AND CompId in ( '0' ")
	for _, v := range compIds {
		sql_build.WriteString(" ,'")
		sql_build.WriteString(v)
		sql_build.WriteString("'")
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		log.Println("FindByMatchIdCompId:", err)
	}
	return dataList
}
