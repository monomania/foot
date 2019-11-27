package entity

import (
	"log"
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	entity2 "tesou.io/platform/foot-parent/foot-core/module/core/entity"
)

//菠菜公司配置
type CompConfig struct {
	//公司ID
	CompId string
	//数据来源信息
	entity2.SourceConfig `xorm:"extends"`
	entity.Base          `xorm:"extends"`
}

func (this *CompConfig) FindByItemId() bool {
	exists, err := mysql.GetEngine().Get(this)
	if err != nil {
		log.Println("FindByItemId:", err)
	}
	return exists
}
