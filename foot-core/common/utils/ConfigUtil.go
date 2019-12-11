package utils

import (
	"github.com/astaxie/beego/config"
	"tesou.io/platform/foot-parent/foot-api/common/base"
)

var (
	//配置信息
	configer config.Configer
)

func init() {
	temp, e := new(config.IniConfig).Parse("conf/app.ini")
	if e != nil {
		base.Log.Info("loadConfig加载配置文件失败:", e.Error())
		return
	}
	configer = temp
}

func GetSection(sectionName string) map[string]string {
	section, e := configer.GetSection(sectionName)
	if e != nil {
		base.Log.Info("未找到对应的配置信息:", sectionName, e.Error())
		return nil
	}
	return section

}
