package vo

type WeekVO struct {
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

	DataList []SuggestVO
}
