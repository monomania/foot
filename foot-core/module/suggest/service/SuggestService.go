package service

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

/**
发布推荐
*/
type SuggestService struct {
	mysql.BaseService
	service.AnalyService
}
