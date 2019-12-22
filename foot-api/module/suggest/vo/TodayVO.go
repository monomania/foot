package vo

type TodayVO struct {
	//开始时间
	BeginDateStr string
	//结束时间
	EndDateStr string
	//数据时间
	DataDateStr string

	DataList []SuggestVO
}
