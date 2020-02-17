package vo

type TTodayVO struct {
	//引入父级属性
	TBaseVO

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

	MainAlflag string
	//红
	MainRedCount int64
	//黑
	MainBlackCount int64
	//胜率
	MainVal string

	DataList []SuggStubVO
}
