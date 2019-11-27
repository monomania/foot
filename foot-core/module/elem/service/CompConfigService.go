package service

import (
	"log"
	"tesou.io/platform/foot-parent/foot-api/module/elem/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//菠菜公司配置
type CompConfigService struct {
	mysql.BaseService
}

func (this *CompConfigService) FindBySId(v *entity.CompConfig) bool {
	exists, err := mysql.GetEngine().Get(v)
	if err != nil {
		log.Println("FindBySId:", err)
	}
	return exists
}
