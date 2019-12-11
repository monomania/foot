package vo

import "encoding/json"

type PriceVO struct {
	//返回值, 0代表成功
	Code int64
	Data []int64
	//msg
	Msg string
}

func (this *PriceVO) ToString() string {
	bytes, _ := json.Marshal(this)
	return string(bytes)
}
