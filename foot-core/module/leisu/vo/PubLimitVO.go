package vo

import "encoding/json"

/**
发布查询到的用户的限制内容
*/
type PubLimitVO struct {
	//返回值, 0代表成功
	Code         int64
	//可推荐次数
	Limit_times  int64
	//可用推荐次数
	Remain_times int64
	//己用的推荐次数
	Used_times   int64
	//msg
	Msg string
}


/**
获取发布次数
*/
func (this *PubLimitVO) ToString() string {
	bytes, _ := json.Marshal(this)
	return string(bytes)
}