package service

import (
	"encoding/json"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	constants2 "tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	utils2 "tesou.io/platform/foot-parent/foot-core/module/leisu/utils"
	vo2 "tesou.io/platform/foot-parent/foot-core/module/leisu/vo"
)

/**
发布前,需要查询限制
 */
type PushLimitService struct {
	//mysql.BaseService
}

/**
获取发布次数
 */
func (this *PushLimitService) GetPublimit() *vo2.PubLimitVO {
	data := utils2.GetText(constants2.PUB_LIMIT_URL)
	if len(data) <= 0 {
		base.Log.Error("GetPublimit:获取到的数据为空")
		return nil
	}
	resp := new(vo2.PubLimitVO)
	json.Unmarshal([]byte(data), resp)
	return resp
}

/**
是否还有发布次数
 */
func (this *PushLimitService) HasPubCount() bool{
	data := this.GetPublimit()
	if data.Remain_times > 0 {
		return true;
	}
	base.Log.Warn(data.ToString())
	return false;
}