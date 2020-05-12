package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
)

type AsiaHisService struct {
	mysql.BaseService
}

func (this *AsiaHisService) Exist(v *pojo.AsiaHis) (string, bool) {
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + v.MatchId + "' AND CompId = '" + v.CompId + "' ")
	temp := &pojo.AsiaHis{}
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
func (this *AsiaHisService) FindByMatchId(matchId string) []*pojo.AsiaHis {
	dataList := make([]*pojo.AsiaHis, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找亚赔
func (this *AsiaHisService) FindByMatchIdCompId(matchId string, compIds ...string) []*pojo.AsiaHis {
	dataList := make([]*pojo.AsiaHis, 0)
	sql_build := strings.Builder{}
	if len(compIds) > 0 && strings.EqualFold(compIds[0], constants.DEFAULT_REFER_ASIA) {
		sql := `
SELECT 
  h.* 
FROM
  foot.t_asia_his h 
WHERE 1 = 1 
  AND h.MatchId = '` + matchId + `' 
  AND h.SLetBall = 
  (SELECT 
    t.SLetBall 
  FROM
    (SELECT 
      h.SLetBall,
      COUNT(1) AS cc 
    FROM
      foot.t_asia_his h 
    WHERE h.MatchId = '` + matchId + `' 
    GROUP BY h.SLetBall) t 
  ORDER BY t.cc DESC 
  LIMIT 0, 1) 
  AND h.ELetBall = (
    (SELECT 
      t.ELetBall 
    FROM
      (SELECT 
        h.ELetBall,
        COUNT(1) AS cc 
      FROM
        foot.t_asia_his h 
      WHERE h.MatchId = '` + matchId + `' 
      GROUP BY h.ELetBall) t 
    ORDER BY t.cc DESC 
    LIMIT 0, 1)
  ) 
ORDER BY h.CompId ASC 
		`
		//执行查询
		this.FindBySQL(sql, &dataList)
	} else {
		sql_build.WriteString(" MatchId = '" + matchId + "' AND CompId in ( '0' ")
		for _, v := range compIds {
			sql_build.WriteString(" ,'")
			sql_build.WriteString(v)
			sql_build.WriteString("'")
		}
		sql_build.WriteString(")")

		err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
		if err != nil {
			base.Log.Error("FindByMatchIdCompId:", err)
		}
	}
	return dataList
}
