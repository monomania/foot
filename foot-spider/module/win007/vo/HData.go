package vo

import "gopkg.in/mgo.v2/bson"

type HData struct {
	//博彩公司名称
	Cn string
	//博彩公司win007id
	CId int
	//初盘3
	Hw float64
	//初盘1
	So float64
	//初盘0
	Gw float64
	//即时盘3
	Rh float64
	//即时盘1
	Rs float64
	//即时盘0
	Rg float64
	//博彩公司级别
	Ct int

	///////////////以上为页面返回////////////////////

	MatchId bson.ObjectId
}
