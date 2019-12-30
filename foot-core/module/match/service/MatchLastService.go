package service

import (
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//足球比赛信息
type MatchLastService struct {
	mysql.BaseService
}

/**
通过比赛时间,主队id,客队id,判断比赛信息是否已经存在
*/
func (this *MatchLastService) FindExists(v *pojo.MatchLast) bool {
	has, err := mysql.GetEngine().Table("`t_match_last`").Where(" `Id` = ?  ", v.Id).Exist()
	if err != nil {
		base.Log.Info("FindExists", err)
	}
	return has
}

/**
查找未结束的比赛
*/
func (this *MatchLastService) FindNotFinished() []*pojo.MatchLast {
	sql_build := `
SELECT 
  la.* 
FROM
  foot.t_match_last la 
  WHERE la.MatchDate > DATE_SUB(NOW(), INTERVAL 2 HOUR)
	`
	//结果值
	dataList := make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build, &dataList)
	return dataList
}

func (this *MatchLastService) FindAll() []*pojo.MatchLast {
	dataList := make([]*pojo.MatchLast, 0)
	mysql.GetEngine().OrderBy("MatchDate").Find(&dataList)
	return dataList
}

/**
查找欧赔不完整的比赛
*/
func (this *MatchLastService) FindEuroIncomplete(count int) []*pojo.MatchLast {
	sql_build := `
SELECT 
  la.* 
FROM
  foot.t_euro_last l,
  foot.t_match_last la 
WHERE l.MatchId = la.Id 
GROUP BY l.MatchId
	`
	sql_build += " HAVING COUNT(1) < " + strconv.Itoa(count)
	//结果值
	dataList := make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build, &dataList)
	return dataList
}
