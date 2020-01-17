package service

import (
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
