package service

import (
	"encoding/json"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	constants2 "tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	utils2 "tesou.io/platform/foot-parent/foot-core/module/leisu/utils"
	vo2 "tesou.io/platform/foot-parent/foot-core/module/leisu/vo"
)

/**
发布前,查询收费
 */
type PriceService struct {
	//mysql.BaseService
}

/**
获取发布次数
 */
func (this *PriceService) GetPrice() *vo2.PriceVO {
	data := utils2.GetText(constants2.PRICE_URL)
	if len(data) <= 0 {
		base.Log.Error("GetPrice:获取到的数据为空")
		return nil
	}
	resp := new(vo2.PriceVO)
	json.Unmarshal([]byte(data), resp)
	return resp
}

/**
是否可以收费
 */
func (this *PriceService) HasPrice() bool{
	data := this.GetPrice()
	if len(data.Data) > 0 {
		base.Log.Warn(data.ToString())
		return true;
	}
	return false;
}