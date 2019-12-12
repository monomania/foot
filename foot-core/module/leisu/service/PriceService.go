package service

import (
	"encoding/json"
	"math/rand"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/module/core/service"
	constants2 "tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	utils2 "tesou.io/platform/foot-parent/foot-core/module/leisu/utils"
	vo2 "tesou.io/platform/foot-parent/foot-core/module/leisu/vo"
)

/**
发布前,查询收费
 */
type PriceService struct {
	service.ConfService
}

/**
获取收费数据
*/
func (this *PriceService) GetPriceVal() int64 {
	var data int64
	entity := this.GetPrice()
	if len(entity.Data) > 0 {
		//如果可以收费,采用收费策略
		var price_strategy string
		config := this.ConfService.GetPubConfig()
		if nil != config {
			price_strategy = config["price_strategy"]
		}
		var index int
		if len(price_strategy) <= 0 {
			//默认不配置,取最大值
			index = len(entity.Data)-1
		}
		//免费
		if strings.EqualFold(price_strategy,"free"){
			return 0
		}
		//最小
		if strings.EqualFold(price_strategy,"min"){
			index = 0
		}
		//最大
		if strings.EqualFold(price_strategy,"max"){
			index = len(entity.Data)-1
		}
		//随机
		if strings.EqualFold(price_strategy,"random"){
			index = rand.Intn(len(entity.Data))
		}
		data = entity.Data[index]
	}else{
		data = 0
	}
	return data
}

/**
获取收费数据
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