package entity

import (
	"log"
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type MatchLastConfig struct {
	/**
	 * 一场比赛一条配置信息
	 */
	MatchConfig `xorm:"extends"`

	MatchId string

	entity.Base `xorm:"extends"`
}

func (this *MatchLastConfig) Query() []*MatchLastConfig {
	params := make([]interface{}, 0)
	sql_build := " 1=1 "
	if this.MatchId != "" {
		sql_build = sql_build + " and MatchId = ? "
		params = append(params, this.MatchId)
	}
	if this.S != "" {
		sql_build = sql_build + " and S = ? "
		params = append(params, this.S)
	}

	entitys := make([]*MatchLastConfig, 0)
	err := mysql.GetEngine().Where(sql_build, params...).Find(&entitys)

	if err != nil {
		log.Println("错误:", err)
	}
	return entitys
}
