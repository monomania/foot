package service

import (
	"log"
	"tesou.io/platform/foot-parent/foot-api/module/elem/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//菠菜公司
type CompService struct {
	mysql.BaseService
}

func (this *CompService) FindExistsByName(v *entity.Comp) bool {
	exist, err := mysql.GetEngine().Get(v)
	if err != nil {
		log.Println("FindExistsByName:", err)
	}
	return exist
}

func (this *CompService) FindAllIds() []string {
	dataList := make([]*entity.Comp, 0)
	dataIdArr := make([]string, 0)
	mysql.GetEngine().Find(&dataList)
	for _, v := range dataList {
		dataIdArr = append(dataIdArr, v.Id)
	}
	return dataIdArr
}
