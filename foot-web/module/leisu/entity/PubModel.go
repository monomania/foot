package entity

type Pub struct {
	//比赛ID
	Id string
	//暂时不明白其意义
	Selects []string
	//所选择的赔率
	Values []string
}

/**
发布推荐
*/
type PubModel struct {
	//标题15字
	Title string
	//内容100字
	Content string
	Price   string
	Multipe string
	//数据信息
	Data Pub
}
