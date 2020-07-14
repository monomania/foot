package pojo

import "tesou.io/platform/foot-parent/foot-api/common/base/pojo"

//每日推荐
type TuijianDaily struct {
	//公众号ID
	AppId string `xorm:" comment('AppId') unique(AppId_Type_MediaId)"`

	Type string `xorm:" comment('Type') unique(AppId_Type_MediaId)"`

	//每日推荐使用的ID
	MediaId string `xorm:" comment('MediaId') unique(AppId_Type_MediaId)"`

	pojo.BasePojo `xorm:"extends"`
}
