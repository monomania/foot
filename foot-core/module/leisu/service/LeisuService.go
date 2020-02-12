package service

import (
	"bytes"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
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
	//过滤重复选项,只取c1
	dataList := make([]*vo2.SuggStubDetailVO, 0)
	for _, e := range tempList {

		if strings.EqualFold(e.AlFlag,"C1"){
			dataList = append(dataList, e)
		}
	}
	//生成内容
	template_paths := []string{
		"assets/common/template/analycontent/001.html",
		"assets/common/template/analycontent/002.html",
		"assets/common/template/analycontent/003.html",
		"assets/common/template/analycontent/004.html",
		"assets/common/template/analycontent/005.html",
		"assets/common/template/analycontent/footer.html",
	}

	for _, e := range dataList {
		var parent_template_path_suffix string
		intn := rand.Intn(23)
		if intn < 10 {
			parent_template_path_suffix = "10" + strconv.Itoa(intn) + ".html"
		} else {
			parent_template_path_suffix = "1" + strconv.Itoa(intn) + ".html"
		}

		parent_template_path := "assets/common/template/analycontent/" + parent_template_path_suffix
		template_paths = append(template_paths, parent_template_path)
		var buf bytes.Buffer
		tpl, err := template.New(parent_template_path_suffix).ParseFiles(template_paths...)
		if err != nil {
			base.Log.Error(err)
		}
		if err := tpl.Execute(&buf, e); err != nil {
			base.Log.Fatal(err)
		}
		e.SContent = buf.String()
		e.SContent = strings.TrimSpace(e.SContent)
		e.SContent = strings.ReplaceAll(e.SContent, " ", "")
		e.SContent = strings.ReplaceAll(e.SContent, "\r", "")
		e.SContent = strings.ReplaceAll(e.SContent, "\n", "")

		base.Log.Info("---------------------------")
		base.Log.Info(parent_template_path_suffix)
		base.Log.Info(e.SContent)
		base.Log.Info("---------------------------")
		e.MatchDateStr = e.MatchDate.Format("02号15:04")
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
