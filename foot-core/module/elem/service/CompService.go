package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//菠菜公司
type CompService struct {
	mysql.BaseService
}

func (this *CompService) Exist(v *pojo.Comp) bool {
	exist, err := mysql.GetEngine().Get(v)
	if err != nil {
		base.Log.Info("ExistByName:", err)
	}
	return exist
}

func (this *CompService) FindAllIds() []string {
	dataList := make([]*pojo.Comp, 0)
	dataIdArr := make([]string, 0)
	mysql.GetEngine().Find(&dataList)
	for _, v := range dataList {
		dataIdArr = append(dataIdArr, v.Id)
	}
	return dataIdArr
}
