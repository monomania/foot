package service

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	pojo2 "tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"time"
)

type C4Service struct {
	AnalyService
	service.BFScoreService
	service.BFBattleService
	service.BFFutureEventService
	service.BFJinService

	//最大让球数据
	MaxLetBall float64
}

func (this *C4Service) ModelName() string {
	return "C4"
}

func (this *C4Service) AnalyTest() {
	this.AnalyService.AnalyTest(this)
}

func (this *C4Service) Analy(analyAll bool) {
	this.AnalyService.Analy(analyAll, this)
}

func (this *C4Service) Analy_Near() {
	this.AnalyService.Analy_Near(this)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *C4Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	if len(temp_data.Id) <= 0 {
		temp_data = new(entity5.AnalyResult)
	}
	matchId := v.Id
	//声明使用变量
	var a18Bet *entity3.AsiaHis
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, constants.DEFAULT_REFER_ASIA)
	if len(aList) < 1 {
		return -1, temp_data
	}
	a18Bet = aList[0]
	temp_data.LetBall = a18Bet.ELetBall
	if math.Abs(a18Bet.ELetBall) > this.MaxLetBall {
		//temp_data.Result =""
		temp_data.PreResult = -1
		return -2, temp_data
	}

	//得出结果
	preResult := -1

	matchDateStr := v.MatchDate.Format("2006-01-02 15:04:05")
	var currentPage, pageSize int64 = 1, 10
	var page *pojo2.Page
	page = new(pojo2.Page)
	page.PageSize = pageSize
	page.CurPage = currentPage
	mainJinList := make([]*pojo.BFJin, 0)
	err1 := this.BFJinService.PageSql("SELECT j.* FROM foot.t_b_f_jin j WHERE j.SclassID = "+v.LeagueId+" AND (j.HomeTeam = '"+v.MainTeamId+"' OR j.GuestTeam = '"+v.MainTeamId+"') AND STR_TO_DATE(j.MatchTimeStr, '%Y%m%d%H%i%s') <  '"+matchDateStr+"' ORDER BY j.MatchTimeStr DESC ", page, &mainJinList)
	if nil != err1 {
		return -2, temp_data
	}

	guestJinList := make([]*pojo.BFJin, 0)
	err2 := this.BFJinService.PageSql("SELECT j.* FROM foot.t_b_f_jin j WHERE j.SclassID = "+v.LeagueId+" AND (j.HomeTeam = '"+v.GuestTeamId+"' OR j.GuestTeam = '"+v.GuestTeamId+"') AND STR_TO_DATE(j.MatchTimeStr, '%Y%m%d%H%i%s') <  '"+matchDateStr+"' ORDER BY j.MatchTimeStr DESC ", page, &guestJinList)
	if nil != err2 {
		return -2, temp_data
	}

	if len(mainJinList) < 10 || len(guestJinList) < 10 {
		return -2, temp_data
	}


	base.Log.Info("比赛时间:", matchDateStr+",对阵:"+v.GuestTeamId, ",初盘让球:", a18Bet.SLetBall, ",即时盘让球:", a18Bet.ELetBall, " ,比分:", v.MainTeamGoals, ":", v.GuestTeamGoals, " ,半场比分:", v.MainTeamHalfGoals, ":", v.GuestTeamHalfGoals)


	mainGoal, guestGoal := 0, 0
	mainScore, guestScore := 0, 0
	for _, temp := range mainJinList {
		if strings.EqualFold(temp.HomeTeam, v.MainTeamId) {
			mainGoal += temp.HomeScore
			if temp.HomeScore > temp.GuestScore {
				mainScore += 3
			}
			if temp.HomeScore == temp.GuestScore {
				mainScore += 1
			}
		}
		if strings.EqualFold(temp.GuestTeam, v.MainTeamId) {
			mainGoal += temp.GuestScore
			if temp.GuestScore > temp.HomeScore {
				mainScore += 3
			}
			if temp.GuestScore == temp.HomeScore {
				mainScore += 1
			}
		}
	}

	for _, temp := range guestJinList {
		if strings.EqualFold(temp.HomeTeam, v.GuestTeamId) {
			guestGoal += temp.HomeScore
			if temp.HomeScore > temp.GuestScore {
				guestScore += 3
			}
			if temp.HomeScore == temp.GuestScore {
				guestScore += 1
			}
		}
		if strings.EqualFold(temp.GuestTeam, v.GuestTeamId) {
			guestGoal += temp.GuestScore
			if temp.GuestScore > temp.HomeScore {
				guestScore += 3
			}
			if temp.GuestScore == temp.HomeScore {
				guestScore += 1
			}
		}
	}

	diffGoal := float64(mainGoal-guestGoal) / 10
	diffScore := float64(mainScore-guestScore) / 10
	eLetBall := a18Bet.ELetBall
	if eLetBall > 0 {
		if diffGoal > 0 || diffScore > 0 {
			if math.Abs(diffGoal-eLetBall) <= 0.25 {
				preResult = 3
			} else if math.Abs(diffScore-eLetBall) <= 0.25 {
				preResult = 3
			} else if math.Abs(diffGoal-eLetBall) >= 0.5 && math.Abs(diffScore-eLetBall) >= 0.5 {
				preResult = 0
			}
		} else {
			preResult = -1
			temp_data.TOVoidDesc = "None"
		}
	}
	if eLetBall < 0 {
		if diffGoal < 0 || diffScore < 0 {
			if math.Abs(diffGoal-eLetBall) <= 0.25 {
				preResult = 0
			} else if math.Abs(diffScore-eLetBall) <= 0.25 {
				preResult = 0
			} else if math.Abs(diffGoal-eLetBall) >= 0.5 && math.Abs(diffScore-eLetBall) >= 0.5 {
				preResult = 3
			}
		} else {
			preResult = -1
			temp_data.TOVoidDesc = "None"
		}
	}

	if v.Id == "1865711" {
		fmt.Println("-------------")
	}

	temp_data.MatchId = v.Id
	temp_data.MatchDate = v.MatchDate
	temp_data.SLetBall = a18Bet.SLetBall
	temp_data.LetBall = a18Bet.ELetBall
	temp_data.Desc = fmt.Sprintf("S:%v,G:%v", diffScore, diffGoal)
	temp_data.AlFlag = this.ModelName()
	format := time.Now().Format("0102150405")
	temp_data.AlSeq = format
	temp_data.PreResult = preResult

	//限制初盘,即时盘让球在0.25以内
	range_letball := math.Abs(a18Bet.SLetBall - a18Bet.ELetBall)
	if (a18Bet.SLetBall > 0 && a18Bet.ELetBall < 0) || (a18Bet.SLetBall < 0 && a18Bet.ELetBall > 0) {
		temp_data.TOVoidDesc = fmt.Sprintf("Turn:%v,%v,%v", strconv.FormatFloat(range_letball, 'f', -1, 64), strconv.FormatFloat(a18Bet.SLetBall, 'f', -1, 64), strconv.FormatFloat(a18Bet.ELetBall, 'f', -1, 64))
		temp_data.PreResult = -1
		temp_data.HitCount = 3
		return -2, temp_data
	}
	if math.Abs(a18Bet.SLetBall-a18Bet.ELetBall) > 0.25 {
		temp_data.TOVoidDesc = fmt.Sprintf("Change:%v,%v,%v", strconv.FormatFloat(range_letball, 'f', -1, 64), strconv.FormatFloat(a18Bet.SLetBall, 'f', -1, 64), strconv.FormatFloat(a18Bet.ELetBall, 'f', -1, 64))
		temp_data.PreResult = -1
		temp_data.HitCount = 3
		return -2, temp_data
	}
	if preResult < 0 {
		temp_data.HitCount = 3
		return -3, temp_data
	}

	if len(temp_data.Id) > 0 {
		temp_data.HitCount = temp_data.HitCount + 1
		//比赛结果
		temp_data.Result = this.IsRight(v, temp_data)
		return 1, temp_data
	} else {
		temp_data.HitCount = 3
		//比赛结果
		temp_data.Result = this.IsRight(v, temp_data)
		return 0, temp_data
	}
}
