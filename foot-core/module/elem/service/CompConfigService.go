package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//菠菜公司配置
type CompConfigService struct {
	mysql.BaseService
}

func (this *CompConfigService) FindBySId(v *pojo.CompConfig) bool {
	exists, err := mysql.GetEngine().Get(v)
	if err != nil {
		base.Log.Info("FindBySId:", err)
	}
	return exists
}
