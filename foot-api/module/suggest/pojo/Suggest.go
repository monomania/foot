package pojo

import (
	pojo2 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
)



/**
发布记录
 */
type Suggest struct {

	pojo2.AnalyResult `xorm:"extends"`
}
