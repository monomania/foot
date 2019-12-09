package vo

import "time"

/**
比赛对应的赔率信息
*/
type OddINFVO struct {
	//类型, 1 胜平负 2 让球胜平负 5北单胜负过关
	DataIdx int
	//赔率
	DataOdd float64
	//11-5北单胜负过关主队 12-5北单胜负过关客队
	DataSelects int
	//提示
	DataTip string
	//让球
	DataPk float64
}

/**
比赛信息
*/
type MatchVO struct {
	DataId          int64
	DataSport       int
	DataMatch       int64
	DataCompetition int
	DataZoomdate    string

	//编号
	Numb string
	//联赛名称
	LeagueName string
	//比赛时间
	MatchDate time.Time
	//主队
	MainTeam string
	//客队
	GuestTeam string
	//赔率选项
	OddDatas []OddINFVO
}

func (this *MatchVO) GetBJDCOddData(option int) *OddINFVO {
	var dataSelects = 11
	if option == 0 {
		dataSelects = 12
	}
	for _, e := range this.OddDatas {
		if e.DataIdx == 5 && e.DataSelects == dataSelects {
			return &e
		}
	}
	return nil
}

func (this *MatchVO) GetOddData(dataIdx int, dataSelects int) *OddINFVO {
	for _, e := range this.OddDatas {
		if e.DataIdx == dataIdx && e.DataSelects == dataSelects {
			return &e
		}
	}
	return nil
}
