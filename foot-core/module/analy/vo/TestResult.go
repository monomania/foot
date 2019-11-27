package vo

import (
	"tesou.io/platform/foot-parent/foot-core/module/match/entity"
	entity2 "tesou.io/platform/foot-parent/foot-core/module/odds/entity")
type AnalyResult struct {
	entity.MatchLast `xorm:"extends"`
	Id               string
	Name             string
	CompName         string
	entity2.EuroLast `xorm:"extends"`
}
