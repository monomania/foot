package vo

type TodayDetailVO struct {
	SpiderDateStr string
	//开始时间
	BeginDateStr string
	//结束时间
	EndDateStr string
	//数据时间
	DataDateStr string

	DataList []SuggestDetailVO

}
