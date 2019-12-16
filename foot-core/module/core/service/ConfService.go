package service

import (
	"strconv"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
)

/**
获取配置信息
*/
type ConfService struct {
}

func (this *ConfService) GetSpiderCycleTime() int64 {
	var result int64
	temp_val := this.GetConfig("spider", "cycle_time")

	if len(temp_val) > 0 {
		result, _ = strconv.ParseInt(temp_val, 0, 64);
	}

	if result <= 0 {
		result = 120
	}
	return result;
}

func (this *ConfService) GetConfig(section string, key string) string {
	var temp_val string
	config := utils.GetSection(section)
	if nil != config {
		temp_val = config[key]
	}
	return temp_val;
}
