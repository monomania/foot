package service

import (
	"encoding/json"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/constants"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/utils"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/vo"
)

/**
发布推荐
*/
type PubService struct {
	service.AnalyService
}

/**
发布比赛
*/
func (this *PubService) Action(param *vo.PubVO) *vo.PubRespVO {
	dataList := this.AnalyService.FindAll()
	for e := range dataList {
		base.Log.Error(e)
	}
	return nil
}

/**
发布比赛
*/
func (this *PubService) Post(param *vo.PubVO) *vo.PubRespVO {
	data := utils.Post(constants.PUB_URL, param)
	if len(data) <= 0 {
		base.Log.Error("Post:获取到的数据为空")
		return nil
	}
	resp := new(vo.PubRespVO)
	json.Unmarshal([]byte(data), resp)
	return resp
}
