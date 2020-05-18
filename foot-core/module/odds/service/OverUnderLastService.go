package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type OverUnderLastService struct {
	mysql.BaseService
}

//查看数据是否已经存在
func (this *OverUnderLastService) Exist(v *pojo.OverUnderLast) (string, bool) {
	temp := &pojo.OverUnderLast{MatchId: v.MatchId, CompId: v.CompId}
	var id string
	exist, err := mysql.GetEngine().Get(temp)
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	if exist {
		id = temp.Id
	}
	return id, exist
}

//根据比赛ID查找亚赔
func (this *OverUnderLastService) FindByMatchId(matchId string) []*pojo.OverUnderLast {
	dataList := make([]*pojo.OverUnderLast, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找亚赔
func (this *OverUnderLastService) FindByMatchIdCompName(matchId string, compNames ...string) []*pojo.OverUnderLast {
	dataList := make([]*pojo.OverUnderLast, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "' AND CompName in ( '0' ")
	for _, v := range compNames {
		sql_build.WriteString(" ,'")
		sql_build.WriteString(v)
		sql_build.WriteString("'")
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchIdCompName:", err)
	}
	return dataList
}
