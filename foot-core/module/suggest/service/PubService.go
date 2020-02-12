package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

/**
发布推荐
*/
type PubService struct {
	mysql.BaseService
}

func (this *PubService) Exist(matchId string) (*pojo.Pub, bool) {
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId)
	temp := &pojo.Pub{}
	exist, err := mysql.GetEngine().Where(sql_build.String()).Get(temp)
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	return temp, exist
}
