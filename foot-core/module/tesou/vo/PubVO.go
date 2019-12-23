package vo

type MatchINFVO struct {
	//比赛ID
	Id int64 `json:"id"`
	//暂时不明白其意义 tr中可获取
	Selects []int `json:"selects"`
	//所选择的赔率
	Values []float64 `json:"values"`
}

/**
发布推荐
*/
type PubVO struct {
	//标题15字
	Title string `json:"title"`
	//内容100字
	Content string `json:"content"`
	//收费价格
	Price   int64  `json:"price"`
	Multiple int64  `json:"multiple"`
	//数据信息
	Data []MatchINFVO `json:"data"`
}
