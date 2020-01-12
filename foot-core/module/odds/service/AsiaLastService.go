package service

import (

	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type AsiaLastService struct {
	mysql.BaseService
}

//查看数据是否已经存在
func (this *AsiaLastService) FindExists(v *pojo.AsiaLast) bool {
	exist, err := mysql.GetEngine().Get(&pojo.AsiaLast{MatchId:v.MatchId,CompId:v.CompId})
	if err != nil {
		base.Log.Error("错误:", err)
	}
	return exist
}

//根据比赛ID查找亚赔
func (this *AsiaLastService) FindByMatchId(matchId string) []*pojo.AsiaLast {
	dataList := make([]*pojo.AsiaLast, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找欧赔
func (this *AsiaLastService) FindByMatchIdCompId(matchId string, compIds ...string) []*pojo.AsiaLast {
	dataList := make([]*pojo.AsiaLast, 0)
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
		base.Log.Error("FindByMatchIdCompId:", err)
	}
	return dataList
}
