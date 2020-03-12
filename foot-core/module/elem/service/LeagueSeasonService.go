package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//联赛表
type LeagueSeasonService struct {
	mysql.BaseService
}

func (this *LeagueSeasonService) Exist(e *pojo.LeagueSeason) (string, bool) {
	sql_build := strings.Builder{}
	sql_build.WriteString(" LeagueId = '" + e.LeagueId + "' AND Season = '" + e.Season + "'")
	temp := &pojo.LeagueSeason{}
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

func (this *LeagueSeasonService) FindByLeagueId(id string) []*pojo.LeagueSeason {
	dataList := make([]*pojo.LeagueSeason, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" LeagueId = '" + id + "'")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByLeagueId:", err)
	}
	return dataList
}
