package entity

/**
发布查询到的用户的限制内容
*/
type PubLimitModel struct {
	//返回值, 0代表成功
	Code         int64
	//可推荐次数
	Limit_times  int64
	//可用推荐次数
	Remain_times int64
	//己用的推荐次数
	Used_times   int64
}
