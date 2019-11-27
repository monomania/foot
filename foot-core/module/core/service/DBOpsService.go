package service

import (
	"encoding/json"
	"log"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	entity4 "tesou.io/platform/foot-parent/foot-core/module/analy/entity"
	entity3 "tesou.io/platform/foot-parent/foot-core/module/elem/entity"
	"tesou.io/platform/foot-parent/foot-core/module/match/entity"
	entity2 "tesou.io/platform/foot-parent/foot-core/module/odds/entity"
)

type DBOpsService struct {
}

/**
 * 清空表
 */
func (this *DBOpsService) TruncateTable(tables []string) {
	engine := mysql.GetEngine()
	for _, v := range tables {
		_, err := engine.Exec(" TRUNCATE TABLE " + v)
		if nil != err {
			log.Println(err)
		}
	}
}

/**
 * xorm支持获取表结构信息，通过调用engine.DBMetas()可以获取到数据库中所有的表，字段，索引的信息。
 */
func (this *DBOpsService) DBMetas() string {
	engine := mysql.GetEngine()
	dbMetas, err := engine.DBMetas()
	if nil != err {
		log.Println(err.Error())
	}
	bytes, _ := json.Marshal(dbMetas)
	result := string(bytes)
	return result
}

/**
 * 同步生成数据库表
 */
func (this *DBOpsService) SyncTableStruct() {
	engine := mysql.GetEngine()
	var err error
	//sync model
	//比赛相关表
	err = engine.Sync2(new(entity.MatchLast),new(entity.MatchLastConfig),new(entity.MatchHis))
	if nil != err {
		log.Println(err.Error())
	}

	//赔率相关表
	err = engine.Sync2(new(entity2.EuroLast),new(entity2.EuroHis),new(entity2.AsiaLast))
	if nil != err {
		log.Println(err.Error())
	}

	//波菜公司，联赛其他数据表
	err = engine.Sync2(new(entity3.Comp),new(entity3.CompConfig),new(entity3.League),new(entity3.LeagueConfig))
	if nil != err {
		log.Println(err.Error())
	}
	//分析的结果集表
	err = engine.Sync2(new(entity4.AnalyResult))
	if nil != err {
		log.Println(err.Error())
	}

}
