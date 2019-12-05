package service

import (
	"encoding/json"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/constants"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/utils"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/vo"
)

/**
发布前,需要查询限制
 */
type PubLimitService struct {
	//mysql.BaseService
}

/**
获取发布次数
 */
func (this *PubLimitService) GetPublimit() *vo.PubLimitVO {
	data := utils.GetText(constants.PUB_LIMIT_URL)
	if len(data) <= 0 {
		base.Log.Error("GetPublimit:获取到的数据为空")
		return nil
	}
	resp := new(vo.PubLimitVO)
	json.Unmarshal([]byte(data), resp)
	return resp
}
