package win007

const MODULE_FLAG = "win007"

/**
 * 比赛分析页面数据
 */
const WIN007_MATCH_ANALY_URL_PATTERN = "http://m.win007.com/analy/Analysis/${matchId}.htm"

/**
 * 欧赔, scheid为比赛Id , cId为公司Id
 */
const WIN007_EUROODD_URL_PATTERN = "http://m.win007.com/Compensate/${matchId}.htm"
const WIN007_EUROODD_BET_URL_PATTERN = "http://m.win007.com/1x2Detail.aspx?scheid=${scheid}&cId=${cId}"

/**
 * 亚赔, scheid为比赛Id , cId为公司Id
 */
const WIN007_ASIAODD_URL_PATTERN = "http://m.win007.com/asian/${matchId}.htm"
const WIN007_ASIAODD_NEW_URL_PATTERN = "http://m.win007.com/HandicapDataInterface.ashx?scheid=${matchId}&type=1&oddskind=0&flesh=0.7215399647784261"



