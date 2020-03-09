package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type EuroTrackService struct {
	mysql.BaseService
}

func (this *EuroTrackService) Exist(v *pojo.EuroTrack) (string, bool) {
	temp := &pojo.EuroTrack{MatchId: v.MatchId, CompId: v.CompId, OddDate: v.OddDate}
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


//根据比赛ID查找欧赔
func (this *EuroTrackService) FindByMatchId(matchId string) []*pojo.EuroTrack {
	dataList := make([]*pojo.EuroTrack, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找欧赔
func (this *EuroTrackService) FindByMatchIdCompId(matchId string, compIds ...string) []*pojo.EuroTrack {
	dataList := make([]*pojo.EuroTrack, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "' AND CompId in ( 0 ")
	for _, v := range compIds {
		sql_build.WriteString(" , ")
		sql_build.WriteString(v)
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).OrderBy("OddDate Desc").Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchIdCompId:", err)
	}
	return dataList
}