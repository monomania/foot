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
	temp_val := utils.GetVal("spider", "cycle_time")

	if len(temp_val) > 0 {
		result, _ = strconv.ParseInt(temp_val, 0, 64);
	}

	if result <= 0 {
		result = 120
	}
	return result;
}
