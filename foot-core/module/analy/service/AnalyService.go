package service

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/analy/vo"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	service3 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"time"
)

type AnalyService struct {
	mysql.BaseService
	service.EuroLastService
	service.EuroHisService
	service.AsiaLastService
	service2.MatchLastService
	service3.LeagueService
	//是否打印赔率数据
	PrintOddData bool
}

func (this *AnalyService) Find(matchId string) *entity5.AnalyResult {
	data := entity5.AnalyResult{MatchId: matchId}
	mysql.GetEngine().Get(&data)
	return &data
}

func (this *AnalyService) FindAll() []*entity5.AnalyResult {
	dataList := make([]*entity5.AnalyResult, 0)
	mysql.GetEngine().OrderBy("CreateTime Desc").Find(&dataList)
	return dataList
}

/**
获取可发布的数据项
1.预算结果是主队
2.比赛未开始
3.比赛未结束
4.alName 算法名称，默认为Euro81_616Service ;
5.option 3(只筛选主队),1(只筛选平局),0(只筛选客队)选项
 */
func (this *AnalyService) GetPubDataList(alName string, option int) []*vo.AnalyResultVO {
	sql_build := strings.Builder{}
	sql_build.WriteString("SELECT ml.`MainTeamId`,ml.`GuestTeamId`,ar.* FROM foot.`t_match_last` ml,foot.`t_analy_result` ar ,(SELECT MAX(temp.`CreateTime`) AS CreateTime FROM foot.`t_analy_result` temp WHERE  temp.`AlFlag` LIKE '" + alName + "%' ) last_analy_time WHERE ml.`Id` = ar.`MatchId` AND ar.`CreateTime` = last_analy_time.CreateTime AND ar.`LeisuPubd` IS FALSE AND ar.`MatchDate` > NOW()   ")
	if option > 0 {
		sql_build.WriteString(" AND ar.`PreResult` = " + strconv.Itoa(option) + " ")
	}
	sql_build.WriteString(" ORDER BY ar.`PreResult` DESC ,ar.`MatchDate` ASC   ")
	//结果值
	entitys := make([]*vo.AnalyResultVO, 0)
	//执行查询
	this.FindBySQL(sql_build.String(), &entitys)
	return entitys
}

//测试加载数据
func (this *AnalyService) LoadData(matchId string) []*entity5.AnalyResult {
	sql_build := strings.Builder{}
	sql_build.WriteString("SELECT ml.*,  bc.id,  bc.name AS compName, el.*  FROM t_match_last ml, t_euro_last el,  t_comp bc  WHERE ml.id = el.matchid  AND el.compid = bc.id AND ml.id =  '" + matchId + "' ")
	//结果值
	entitys := make([]*entity5.AnalyResult, 0)
	//执行查询
	this.FindBySQL(sql_build.String(), &entitys)
	return entitys
}

/**
比赛的实际结果计算
 */
func (this *AnalyService) ActualResult(last *entity3.AsiaLast, v *entity2.MatchLast) int {
	var result int
	h2, _ := time.ParseDuration("2h")
	matchDate := v.MatchDate.Add(h2)
	if matchDate.After(time.Now()) {
		//比赛未结束
		return -1
	}

	elb_sum := last.ELetBall
	var mainTeamGoals float64
	if elb_sum > 0 {
		mainTeamGoals = float64(v.MainTeamGoals) - elb_sum
	} else {
		mainTeamGoals = float64(v.MainTeamGoals) + math.Abs(elb_sum)
	}
	diff_goals := float64((v.MainTeamGoals - v.GuestTeamGoals)) - elb_sum
	if diff_goals <= 0.25 && diff_goals >= -0.25 {
		result = 1
	} else if mainTeamGoals > float64(v.GuestTeamGoals) {
		result = 3
	} else if mainTeamGoals < float64(v.GuestTeamGoals) {
		result = 0
	} else {
		result = 1
	}
	return result
}

/**
1.欧赔是主降还是主升 主降为true
*/
func EuroMainDown(e81data *entity3.EuroLast, e616data *entity3.EuroLast) int {
	if e81data.Ep3 < e81data.Sp3 && e616data.Ep3 < e616data.Sp3 {
		return 3
	} else if e81data.Ep0 < e81data.Sp0 && e616data.Ep0 < e616data.Sp0 {
		return 0
	}
	return 1
}

/**
2.亚赔是主降还是主升 主降为true
*/
func AsiaMainDown(a18betData *entity3.AsiaLast) bool {
	slb_sum := a18betData.SLetBall
	elb_sum := a18betData.ELetBall

	if elb_sum > slb_sum {
		return true
	} else if elb_sum < slb_sum {
		return false
	} else { //初始让球和即时让球一致
		if a18betData.Ep3 < a18betData.Sp3 {
			return true
		}
	}
	return false
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
