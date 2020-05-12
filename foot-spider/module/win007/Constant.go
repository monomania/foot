package win007

const (
	MODULE_FLAG = "win007"

	WIN007_BASE_URL = "http://m.win007.com/"

	/**
	 * 比赛分析页面数据
	 */
	WIN007_BASE_FACE_URL_PATTERN = "http://m.win007.com/analy/Analysis/${matchId}.htm"

	/**
	 * 欧赔, scheid为比赛Id , cId为公司Id
	 */
	WIN007_EUROODD_URL_PATTERN     = "http://m.win007.com/Compensate/${matchId}.htm"
	WIN007_EUROODD_BET_URL_PATTERN = "http://m.win007.com/1x2Detail.aspx?scheid=${scheid}&cId=${cId}"

	/**
	 * 亚赔, scheid为比赛Id , cId为公司Id
	 */
	WIN007_ASIAODD_URL_PATTERN     = "http://m.win007.com/asian/${matchId}.htm"
	WIN007_ASIAODD_NEW_URL_PATTERN = "http://m.win007.com/HandicapDataInterface.ashx?scheid=${matchId}&type=1&oddskind=0&flesh=${flesh}"

	/**
	资料库里的比赛赛程前缀
	示例:http://m.win007.com/info/Fixture/2019-2020/34_0_0.htm
	 */
	WIN007_MATCH_HIS_PATTERN = "http://m.win007.com/info/Fixture/${season}/${leagueId}_${subId}_${round}.htm"
)

const (
	SLEEP_RAND_S = 10
	SLEEP_RAND_E = 100

)
