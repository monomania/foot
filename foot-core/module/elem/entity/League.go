package entity

import (
	"log"
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//联赛表
type League struct {
	//联赛名称
	Name string

	//联赛级别
	Level int

	//联赛官网
	Website string

	entity.Base `xorm:"extends"`
}

/**
 */
func (this *League) FindExistsByName() bool {
	exist, err := mysql.GetEngine().Exist(&League{Name: this.Name})
	if err != nil {
		log.Println("错误:", err)
	}
	return exist
}

func (this *League) FindExistsById() bool {
	league := new(League)
	league.Id = this.Id
	exist, err := mysql.GetEngine().Exist(league)
	if err != nil {
		log.Println("错误:", err)
	}
	return exist
}

func (this *League) FindById() {
	_, err := mysql.GetEngine().Get(this)
	if err != nil {
		log.Println("错误:", err)
	}
}
