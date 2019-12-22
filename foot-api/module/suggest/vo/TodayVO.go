package vo

type TodayVO struct {
	//开始时间
	BeginDateStr string
	//结束时间
	EndDateStr string

	DataList []SuggestVO
}
