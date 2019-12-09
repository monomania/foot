package vo

import "encoding/json"

/**
发布推荐结果
*/
type PubRespVO struct {
	//返回值, 0代表成功
	Code int64
	//ID
	Id int64
	//msg
	Msg string
}


/**
获取发布次数
 */
func (this *PubRespVO) ToString() string {
	bytes, _ := json.Marshal(this)
	return string(bytes)
}
