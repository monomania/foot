package service

import (
	"strconv"
	"strings"
	vo2 "tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	"tesou.io/platform/foot-parent/foot-core/module/suggest/service"
	"time"
)

type LeisuService struct {
	service.SuggestService
}

/**
获取可发布的数据
 */
func (this *LeisuService) ListPubAbleData() []*vo2.SuggStubDetailVO {
	//获取分析计算出的比赛列表
	param := new(vo2.SuggStubDetailVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-1h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	//只推送稳胆
	tempList := this.SuggestService.QueryLeisu(param)
	if len(tempList) <= 0 {
		return tempList
	}
	//过滤重复选项
	dataList := make([]*vo2.SuggStubDetailVO, 0)
	for _, e := range tempList {
		for _, e2 := range dataList {
			if strings.EqualFold(e.MatchId, e2.MatchId) {
				continue
			}
			dataList = append(dataList, e)
		}
	}
	return dataList
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
