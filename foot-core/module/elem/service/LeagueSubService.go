package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//联赛表
type LeagueSubService struct {
	mysql.BaseService
}

func (this *LeagueSubService) Exist(e *pojo.LeagueSub) (string, bool) {
	sql_build := strings.Builder{}
	sql_build.WriteString(" LeagueId = '" + e.LeagueId + "' AND Season = '" + e.Season + "' AND SubId = '" + e.SubId + "'")
	temp := &pojo.LeagueSub{}
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

func (this *LeagueSubService) FindByLeagueId(id string) []*pojo.LeagueSub {
	dataList := make([]*pojo.LeagueSub, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" LeagueId = '" + id + "'")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByLeagueId:", err)
	}
	return dataList
}
