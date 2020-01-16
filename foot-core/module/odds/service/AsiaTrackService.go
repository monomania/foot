package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type AsiaTrackService struct {
	mysql.BaseService
}

func (this *AsiaTrackService) FindExists(v *pojo.AsiaTrack) bool {
	exist, err := mysql.GetEngine().Get(&pojo.AsiaTrack{MatchId:v.MatchId,CompId:v.CompId,OddDate:v.OddDate})
	if err != nil {
		base.Log.Error("FindExists:", err)
	}
	return exist
}

//根据比赛ID查找亚赔
func (this *AsiaTrackService) FindByMatchId(matchId string) []*pojo.AsiaTrack {
	dataList := make([]*pojo.AsiaTrack, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找亚赔
func (this *AsiaTrackService) FindByMatchIdCompId(matchId string, compIds ...string) []*pojo.AsiaTrack {
	dataList := make([]*pojo.AsiaTrack, 0)
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
