package vo

type MonthVO struct {
	//开始时间
	BeginDateStr string
	//结束时间
	EndDateStr string

	//总场次
	MatchCount int64
	//红
	RedCount int64
	//走
	WalkCount int64
	//黑
	BlackCount int64
	//最长连红
	LinkRedCount int64
	//最长连黑
	LinkBlackCount int64
	//胜率
	Val string

	//红
	MainRedCount int64
	//黑
	MainBlackCount int64
	//胜率
	MainVal string

	DataList []SuggestVO
}
