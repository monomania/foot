package controller

import (
	"github.com/astaxie/beego"
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type BaseController struct {
	beego.Controller
	mysql.BaseService
}



/**
获取客户端ip地址
*/
func (this *BaseController) GetIp() string {
	return ""
}

func (this *BaseController) newPage() *entity.Page {
	page := new(entity.Page)
	page.SetPageSize(10)
	page.SetCurPage(1)
	return page
}

/**
获取分页设置
 */
func (this *BaseController) GetPage() *entity.Page {
	pageSize := this.Input().Get("pageSize")
	curPage := this.Input().Get("curPage")
	orderType := this.Input().Get("orderType")
	orderName := this.Input().Get("orderName")

	page := this.newPage()
	if len(pageSize) > 0 {
		i, _ := strconv.Atoi(pageSize)
		page.SetPageSize(int64(i))
	}

	if len(curPage) > 0 {
		i, _ := strconv.Atoi(curPage)
		page.SetCurPage(int64(i))
	}

	if orderType != "" {
		page.SetOrderType(orderType)
	}

	if orderName != "" {
		page.SetOrderName(orderName)
	}
	return page
}



