package service

import "tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"

/**
发布前,需要查询限制
 */
type PubLimitService struct {
	mysql.BaseService
}
