package service

import (
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
)

type OverUnderHisService struct {
	mysql.BaseService
}

func (this *OverUnderHisService) Exist(v *pojo.OverUnderHis) (string, bool) {
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + v.MatchId + "' AND CompId = '" + strconv.Itoa(v.CompId) + "' ")
	temp := &pojo.OverUnderHis{}
	var id string
	exist, err := mysql.GetEngine().Where(sql_build.String()).Get(temp)
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	if exist {
		id = temp.Id
	}
	return id, exist
}

//根据比赛ID查找亚赔
func (this *OverUnderHisService) FindByMatchId(matchId string) []*pojo.OverUnderHis {
	dataList := make([]*pojo.OverUnderHis, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找亚赔
func (this *OverUnderHisService) FindByMatchIdCompName(matchId string, compNames ...string) []*pojo.OverUnderHis {
	dataList := make([]*pojo.OverUnderHis, 0)
	sql_build := strings.Builder{}
	if len(compNames) > 0 && strings.EqualFold(compNames[0], constants.DEFAULT_REFER_ASIA) {
		sql := `
SELECT 
  h.* 
FROM
  t_asia_his h 
WHERE 1 = 1 
  AND h.MatchId = '` + matchId + `' 
  AND h.PanKou = 
  (SELECT 
    t.PanKou 
  FROM
    (SELECT 
      h.PanKou,
      COUNT(1) AS cc 
    FROM
      t_over_under_his h 
    WHERE h.MatchId = '` + matchId + `' 
    GROUP BY h.PanKou) t 
  ORDER BY t.cc DESC 
  LIMIT 0, 1) 
  AND h.EPanKou = (
    (SELECT 
      t.EPanKou 
    FROM
      (SELECT 
        h.EPanKou,
        COUNT(1) AS cc 
      FROM
        t_over_under_his h 
      WHERE h.MatchId = '` + matchId + `' 
      GROUP BY h.EPanKou) t 
    ORDER BY t.cc DESC 
    LIMIT 0, 1)
  ) 
ORDER BY h.CompName ASC 
		`
		//执行查询
		this.FindBySQL(sql, &dataList)
	} else {
		sql_build.WriteString(" MatchId = '" + matchId + "' AND CompName in ( '0' ")
		for _, v := range compNames {
			sql_build.WriteString(" ,'")
			sql_build.WriteString(v)
			sql_build.WriteString("'")
		}
		sql_build.WriteString(")")

		err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
		if err != nil {
			base.Log.Error("FindByMatchIdCompName:", err)
		}
	}
	return dataList
}
