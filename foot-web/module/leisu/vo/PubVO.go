package vo

type MatchINFVO struct {
	//比赛ID
	Id string
	//暂时不明白其意义 tr中可获取
	Selects []string
	//所选择的赔率
	Values []string
}

/**
发布推荐
*/
type PubVO struct {
	//标题15字
	Title string
	//内容100字
	Content string
	Price   string
	Multipe string
	//数据信息
	Data MatchINFVO
}
