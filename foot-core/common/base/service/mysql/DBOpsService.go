package mysql

import (
	"encoding/json"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity4 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	entity1 "tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/pojo"
)

type DBOpsService struct {
}

/**
 * 清空表
 */
func (this *DBOpsService) TruncateTable(tables []string) {
	engine := GetEngine()
	for _, v := range tables {
		_, err := engine.Exec(" TRUNCATE TABLE " + v)
		if nil != err {
			base.Log.Error(err)
		}
	}
}

/**
 * xorm支持获取表结构信息，通过调用engine.DBMetas()可以获取到数据库中所有的表，字段，索引的信息。
 */
func (this *DBOpsService) DBMetas() string {
	engine := GetEngine()
	dbMetas, err := engine.DBMetas()
	if nil != err {
		base.Log.Error(err.Error())
	}
	bytes, _ := json.Marshal(dbMetas)
	result := string(bytes)
	return result
}

/**
 * 同步生成数据库表
 */
func (this *DBOpsService) SyncTableStruct() {
	engine := GetEngine()
	var err error
	//sync model
	//比赛相关表
	err = engine.Sync2(new(entity1.MatchLast), new(entity1.MatchHis))
	if nil != err {
		base.Log.Error(err.Error())
	}

	//赔率相关表
	err = engine.Sync2(new(entity2.EuroLast), new(entity2.EuroHis), new(entity2.AsiaLast), new(entity2.AsiaHis))
	if nil != err {
		base.Log.Error(err.Error())
	}

	//波菜公司，联赛其他数据表
	err = engine.Sync2(new(entity3.Comp), new(entity3.League))
	if nil != err {
		base.Log.Error(err.Error())
	}
	//分析的结果集表
	err = engine.Sync2(new(entity4.AnalyResult))
	if nil != err {
		base.Log.Error(err.Error())
	}

	//发布相关的数据库表
	err = engine.Sync2(new(pojo.Pub))
	if nil != err {
		base.Log.Error(err.Error())
	}
}
