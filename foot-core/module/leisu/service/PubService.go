package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	vo3 "tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	constants2 "tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	utils2 "tesou.io/platform/foot-parent/foot-core/module/leisu/utils"
	vo2 "tesou.io/platform/foot-parent/foot-core/module/leisu/vo"
	"time"
)

/**
发布推荐
*/
type PubService struct {
	LeisuService
	MatchPoolService
	PubLimitService
	PriceService
}

/**
获取周期间隔时间
*/
func (this *PubService) CycleTime() int64 {
	var result int64
	temp_val := utils.GetVal(constants2.SECTION_NAME, "cycle_time")
	if len(temp_val) > 0 {
		result, _ = strconv.ParseInt(temp_val, 0, 64);
	}

	if result <= 0 {
		result = 120
	}
	return result
}

/**
发布北京单场胜负过关
*/
func (this *PubService) PubBJDC() {
	//获取分析计算出的比赛列表
	tempList := this.LeisuService.ListPubAbleData()
	if len(tempList) < 1 {
		base.Log.Info(fmt.Sprintf("1.当前没有可发布的比赛!!!!"))
		return
	}

	//获取发布池的比赛列表
	matchPool := this.MatchPoolService.GetMatchList()
	//适配比赛,获取发布列表
	pubList := make(map[*vo3.SuggStubDetailVO]*vo2.MatchVO, 0)
	for _, temp := range tempList {
		analy_mainTeam := temp.MainTeam
		analy_guestTeam := temp.GuestTeam
		for _, match := range matchPool {
			if strings.EqualFold(analy_mainTeam, match.MainTeam) {
				pubList[temp] = match
				break;
			}
			if strings.EqualFold(analy_guestTeam, match.GuestTeam) {
				pubList[temp] = match
				break;
			}
		}
	}

	if len(pubList) <= 0 {
		base.Log.Info(fmt.Sprintf("2.当前无可发布的比赛!!!!比赛池size:%d,分析赛果size:%d", len(matchPool), len(tempList)))
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
				//this.AnalyService.Modify(&analy.AnalyResult)
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
	var title string
	matchDate := param.MatchDate.Format("20060102150405")
	titleTpl := utils.GetVal(constants2.SECTION_NAME, "title_tpl")
	if len(titleTpl) > 0 {
		titleTpl = strings.ReplaceAll(titleTpl, "{leagueName}", param.LeagueName)
		titleTpl = strings.ReplaceAll(titleTpl, "{matchDate}", matchDate)
		titleTpl = strings.ReplaceAll(titleTpl, "{mainTeam}", param.MainTeam)
		titleTpl = strings.ReplaceAll(titleTpl, "{guestTeam}", param.GuestTeam)
		title = titleTpl
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
const pubContent = `
本次推荐为程序全自动处理,全程无人为参与干预.
进而避免了人为分析的主观性及不稳定因素.
程序根据各大波菜多维度数据,结合作者多年足球分析经验,十年程序员生涯,精雕细琢历经26个月得出的产物.
程序执行流程包括且不仅限于(数据自动获取-->分析学习-->自动推送发布).
经近三个月的实验准确率一直能维持在一个较高的水平.
依据该项目为依托已经吸引了不少朋友,现目前通过雷速号再次验证程序的准确率,望大家长期关注校验.!
`
func (this *PubService) content(param *vo2.MatchVO) string {
	var content string
	matchDate := param.MatchDate.Format("20060102150405")
	contentTpl := utils.GetVal(constants2.SECTION_NAME, "content_tpl")
	if len(contentTpl) > 0 {
		contentTpl = strings.ReplaceAll(contentTpl, "{leagueName}", param.LeagueName)
		contentTpl = strings.ReplaceAll(contentTpl, "{matchDate}", matchDate)
		contentTpl = strings.ReplaceAll(contentTpl, "{mainTeam}", param.MainTeam)
		contentTpl = strings.ReplaceAll(contentTpl, "{guestTeam}", param.GuestTeam)
		content = contentTpl
	}

	if len(contentTpl) <= (2 * 101) {
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
	pubVO.Content = this.content(param)
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
