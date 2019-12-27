package service

import (
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/analy/vo"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	"time"
)


type LeisuService struct {
	service.AnalyService
}



func (this *LeisuService) ListDefaultData() []*vo.AnalyResultVO {
	teamOption := this.teamOption()
	al_flag := utils.GetVal(constants.SECTION_NAME, "al_flag")
	hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
	hit_count, _ := strconv.Atoi(hit_count_str)
	//获取分析计算出的比赛列表
	analyList := this.AnalyService.ListData(al_flag, hit_count, teamOption)
	return analyList
}


/**
###推送的主客队选项,
#格式为:时间:选项,时间:选项,时间:选项
#时间只支持设置小时数
#3 只推送主队, 1 只推送平局, 0 只推送客队,-1 全部推送
#示例0-3:-1,4-19:3,19-23:-1,未设置时间段为默认只推送3
*/
func (this *LeisuService) teamOption() int {
	var result int
	tempOptionConfig := utils.GetVal(constants.SECTION_NAME, "team_option")
	if len(tempOptionConfig) <= 0 {
		//默认返回 主队选项
		return 3
	}
	//当前的小时
	currentHour, _ := strconv.Atoi(time.Now().Format("15"))
	hourRange_options := strings.Split(tempOptionConfig, ",")
	for _, e := range hourRange_options {
		h_o := strings.Split(e, ":")
		hourRanges := strings.Split(h_o[0], "-")
		option, _ := strconv.Atoi(h_o[1])
		hourBegin, _ := strconv.Atoi(hourRanges[0])
		hourEnd, _ := strconv.Atoi(hourRanges[1])
		if hourBegin <= currentHour && currentHour <= hourEnd {
			result = option
			break;
		}
	}
	return result
}