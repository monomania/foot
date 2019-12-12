package service

import (
	"tesou.io/platform/foot-parent/foot-core/common/utils"
)

/**
获取配置信息
*/
type ConfService struct {

}

func (this *ConfService) GetPubConfig() map[string]string {
	section := utils.GetSection("pub")
	return section
}


