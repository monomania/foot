package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type EuroHisService struct {
	mysql.BaseService
}

func (this *EuroHisService) Exist(v *pojo.EuroHis) (string, bool) {
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + v.MatchId + "' AND CompId = '" + v.CompId + "' ")
	temp := &pojo.EuroHis{}
	var id string
	exist, err := mysql.GetEngine().Where(sql_build.String()).Get(temp)
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	if exist {
		id = temp.Id
	}
	return id, exist
}

//根据比赛ID查找欧赔
func (this *EuroHisService) FindByMatchId(matchId string) []*pojo.EuroHis {
	dataList := make([]*pojo.EuroHis, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找欧赔
func (this *EuroHisService) FindByMatchIdCompId(matchId string, compIds ...string) []*pojo.EuroHis {
	dataList := make([]*pojo.EuroHis, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "' AND CompId in ( 0 ")
	for _, v := range compIds {
		sql_build.WriteString(" , ")
		sql_build.WriteString(v)
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchIdCompId:", err)
	}
	return dataList
}
