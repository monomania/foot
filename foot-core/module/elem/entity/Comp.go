package entity

import (
	"log"
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//菠菜公司
type Comp struct {
	//名称
	Name string
	//公司网址
	Wesite string

	entity.Base `xorm:"extends"`
}

func (this *Comp) FindExistsByName() bool {
	exist, err := mysql.GetEngine().Get(this)
	if err != nil {
		log.Println("错误:", err)
	}
	return exist
}

func (this *Comp) FindAllIds() []string {
	dataList := make([]*Comp, 0)
	dataIdArr := make([]string, 0)
	mysql.GetEngine().Find(&dataList)
	for _, v := range dataList {
		dataIdArr = append(dataIdArr, v.Id)
	}
	return dataIdArr
}
