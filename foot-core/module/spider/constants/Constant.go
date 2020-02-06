package constants

import "time"

var (
	//数据爬取时间
	SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")
	//全量数据爬取时间
	FullSpiderDateStr = time.Now().Format("2006-01-02 15:04:05")
)
