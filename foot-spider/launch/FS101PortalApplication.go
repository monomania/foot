package launch

func Spider(matchLevel int) {
	Before_spider_match()
	Before_spider_asiaLast()
	Before_spider_euroLast()
	Before_spider_euroHis()
	//执行抓取比赛数据
	//执行抓取比赛欧赔数据
	//执行抓取亚赔数据
	//执行抓取欧赔历史
	Spider_match(matchLevel)
	Spider_asiaLast()
	Spider_euroLast()
	Spider_euroHis()
}
