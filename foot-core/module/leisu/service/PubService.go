package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/analy/vo"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/core/service"
	constants2 "tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	utils2 "tesou.io/platform/foot-parent/foot-core/module/leisu/utils"
	vo2 "tesou.io/platform/foot-parent/foot-core/module/leisu/vo"
	"time"
)

/**
发布推荐
*/
type PubService struct {
	service2.ConfService
	service.AnalyService
	MatchPoolService
	PubLimitService
	PriceService
}

func (this *PubService) getConfig(key string) string {
	var temp_val string
	config := this.ConfService.GetPubConfig()
	if nil != config {
		temp_val = config[key]
	}
	return temp_val;
}

/**
获取周期间隔时间
*/
func (this *PubService) CycleTime() int64 {
	var result int64
	temp_val := this.getConfig("cycle_time")

	if len(temp_val) > 0 {
		result, _ = strconv.ParseInt(temp_val, 0, 64);
	}

	if result <= 0 {
		result = 186
	}
	return result
}

/**
###推送的主客队选项,
#格式为:时间:选项,时间:选项,时间:选项
#时间只支持设置小时数
#3 只推送主队, 1 只推送平局, 0 只推送客队,-1 全部推送
#示例0-3:-1,4-19:3,19-23:-1,未设置时间段为默认只推送3
*/
func (this *PubService) teamOption() int {
	var result int
	tempOptionConfig := this.getConfig("team_option")
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

/**
发布北京单场胜负过关
*/
func (this *PubService) PubBJDC() {
	teamOption := this.teamOption()
	al_flag := this.getConfig("al_flag")
	hit_count, _ := strconv.Atoi(this.getConfig("hit_count"))
	//获取分析计算出的比赛列表
	analyList := this.AnalyService.GetPubDataList(al_flag, hit_count, teamOption)
	if len(analyList) < 1 {
		base.Log.Info(fmt.Sprintf("1.当前没有可发布的比赛,发布的TeamOption为%d!!!!", teamOption))
		return
	}

	//获取发布池的比赛列表
	matchPool := this.MatchPoolService.GetMatchList()
	//适配比赛,获取发布列表
	pubList := make(map[*vo.AnalyResultVO]*vo2.MatchVO, 0)
	for _, analy := range analyList {
		analy_mainTeam := analy.MainTeamId
		for _, match := range matchPool {
			if strings.EqualFold(analy_mainTeam, match.MainTeam) {
				pubList[analy] = match
				break;
			}
		}
	}

	if len(pubList) <= 0 {
		base.Log.Info(fmt.Sprintf("2.当前无可发布的比赛!!!!比赛池size:%d,分析赛果size:%d", len(matchPool), len(analyList)))
		return
	}
	//打印要发布的比赛
	for _, match := range pubList {
		base.Log.Info(fmt.Sprintf("3.即将发布的比赛:%v%v%v%vVS%v", match.MatchDate, match.Numb, match.LeagueName, match.MainTeam, match.GuestTeam))
	}

	//发布比赛
	for analy, match := range pubList {
		//检查是否还有发布次数
		count := this.PubLimitService.HasPubCount()
		if !count {
			base.Log.Error("已经没有发布次数啦,请更换cookies帐号!")
			hours, _ := strconv.Atoi(time.Now().Format("15"))
			time.Sleep(time.Duration(int64(24-hours)) * time.Hour)
			break;
		}
		//检查是否是重复发布

		//------------------------------------------------
		action := this.BJDCAction(match, analy.PreResult)
		if nil != action {
			switch action.Code {
			case 0, 100002:
				//0 成功 100002 每场比赛同一种玩法只可选择1次
				analy.LeisuPubd = true
				this.AnalyService.Modify(&analy.AnalyResult)
			case 100003:
				//100003 标题长度不正确
			default:
				break
			}
		}
		//需要至少间隔5分钟,最大14分钟，再进行下一次发布
		interval := rand.Intn(10) + 5
		base.Log.Info("随机间隔时间为:", interval)
		time.Sleep(time.Duration(interval) * time.Minute)
	}
}

/**
获取标题
*/
func (this *PubService) title(param *vo2.MatchVO) string {
	matchDate := param.MatchDate.Format("20060102150405")

	var title string
	var titleTemplate string
	config := this.ConfService.GetPubConfig()
	if nil != config {
		titleTemplate = config["title"]
	}

	if len(titleTemplate) > 0 {
		titleTemplate = strings.ReplaceAll(titleTemplate, "{leagueName}", param.LeagueName)
		titleTemplate = strings.ReplaceAll(titleTemplate, "{matchDate}", matchDate)
		titleTemplate = strings.ReplaceAll(titleTemplate, "{mainTeam}", param.MainTeam)
		titleTemplate = strings.ReplaceAll(titleTemplate, "{guestTeam}", param.GuestTeam)
		title = titleTemplate
	} else {
		title = param.LeagueName + " " + matchDate + " " + param.MainTeam + "VS" + param.GuestTeam
	}
	if len(title) < (2 * 16) {
		title = "大数据AI精推:" + title
	}
	return title
}

/**
获取内容
*/
const pubContent = "本次推荐为程序全自动处理,全程无人为参与干预.进而避免了人为分析的主观性及不稳定因素.程序根据各大波菜多维度数据,结合作者多年足球分析经验,十年程序员生涯,精雕细琢历经26个月得出的产物.程序执行流程包括且不仅限于(数据自动获取-->分析学习-->自动推送发布).经近三个月的实验准确率一直能维持在一个较高的水平.依据该项目为依托已经吸引了不少朋友,现目前通过雷速号再次验证程序的准确率,望大家长期关注校验.!"

func (this *PubService) content() string {
	var content string
	config := this.ConfService.GetPubConfig()
	if nil != config {
		content = config["content"]
	}
	if len(content) <= (2 * 101) {
		content = pubContent
	}
	return content
}

/**
发布比赛
*/
func (this *PubService) BJDCAction(param *vo2.MatchVO, option int) *vo2.PubRespVO {
	pubVO := new(vo2.PubVO)
	pubVO.Title = this.title(param)
	pubVO.Content = this.content()
	pubVO.Multiple = 0
	pubVO.Price = this.PriceService.GetPriceVal()
	//设置赔率
	oddData := param.GetBJDCOddData(option)
	if oddData == nil {
		//没有找到北单胜负过关的选项
		return nil;
	}
	infvo := vo2.MatchINFVO{}
	infvo.Id = param.DataId
	infvo.Selects = []int{oddData.DataSelects}
	infvo.Values = []float64{oddData.DataOdd}
	pubVO.Data = []vo2.MatchINFVO{infvo}
	//执行发布
	post := this.PubPost(pubVO)
	base.Log.Info("发布结果:" + pubVO.Title + " " + post.ToString())
	return post
}

/**
处理http post
*/
func (this *PubService) PubPost(param *vo2.PubVO) *vo2.PubRespVO {
	data := utils2.Post(constants2.PUB_URL, param)
	if len(data) <= 0 {
		base.Log.Error("PubPost:获取到的数据为空")
		return nil
	}
	//base.Log.Info("http post 请求返回:" + data)
	resp := new(vo2.PubRespVO)
	json.Unmarshal([]byte(data), resp)
	return resp
}
