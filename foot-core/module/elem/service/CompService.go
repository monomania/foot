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

func (this *CompService) FindEuroIds() []string {
	sql_build_1 := ` 
		SELECT tc.* FROM foot.t_comp tc WHERE tc.type = 1 
	`
	//结果值
	dataList := make([]*pojo.Comp, 0)
	//执行查询
	this.FindBySQL(sql_build_1, &dataList)

	dataIdArr := make([]string, 0)
	for _, v := range dataList {
		dataIdArr = append(dataIdArr, v.Id)
	}
	return dataIdArr
}
