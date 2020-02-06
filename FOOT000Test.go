package main

import (
	"fmt"
	"strconv"
	"tesou.io/platform/foot-parent/foot-core/module/spider/constants"
	"time"
)

func abort(funcname string, err string) {
	panic(funcname + " failed: " + err)
}

func print_version(v uint32) {
	major := byte(v)
	minor := uint8(v >> 8)
	build := uint16(v >> 16)
	print("windows version ", major, ".", minor, " (Build ", build, ")\n")
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}


var aaa int

func set()  {
	constants.SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")
	constants.FullSpiderDateStr = constants.SpiderDateStr
}


func main(){
	h2, _ := time.ParseDuration("129m")
	//比赛结束的时间点
	matchEndDate := time.Now().Add(h2)
	fmt.Println(matchEndDate.Format("2006-01-02 15:04:05"))
}

//func main() {
//	h, err := syscall.LoadLibrary("kernel32.dll")
//	if err != nil {
//		abort("LoadLibrary", err.Error())
//	}
//	defer syscall.FreeLibrary(h)
//	proc, err := syscall.GetProcAddress(h, "GetVersion")
//	if err != nil {
//		abort("GetProcAddress", err.Error())
//	}
//	r, _, _ := syscall.Syscall(uintptr(proc), 0, 0, 0, 0)
//	print_version(uint32(r))
//
//
//
//}

/*
func main() {
	var mainRedCount, mainBlackCount int64
	mainRedCount =27
	mainBlackCount = 22
	value := float64(mainRedCount) / (float64(mainRedCount) + float64(mainBlackCount)) *100
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	fmt.Println(strconv.FormatFloat(value,'f', -1, 64)+"%")

	//测试
	suggestService := new(service3.SuggestService)
	i := vo.SuggStubVO{}
	//i.AlFlag = "Euro20191212Service"
	i.BeginDateStr = "2019-12-19 00:00:00"
	query := suggestService.Query(&i)
	for i, e := range query {
		bytes, _ := json.Marshal(e)
		fmt.Println(fmt.Sprintf("%d,%v", i, string(bytes)))
	}

	//测试markdown
	tpl, err := template.New("week.html").Funcs(getFuncMap()).ParseFiles("assets/wechat/html/week.html")
	if err != nil {
		base.Log.Error(err)
	}
	weekVO := vo.TWeekVO{}
	weekVO.MatchCount = 98
	weekVO.RedCount = 68
	weekVO.WalkCount = 40
	weekVO.BlackCount = 20
	weekVO.LinkRedCount = 10
	weekVO.LinkBlackCount = 5
	dataList := make([]vo.SuggStubVO, 1)
	suggestVO := vo.SuggStubVO{}
	suggestVO.LeagueName = "联赛1"
	suggestVO.MatchDate = time.Now()
	suggestVO.MainTeam = "主队1"
	dataList[0] = suggestVO
	weekVO.DataList = dataList
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, &weekVO); err != nil {
		base.Log.Fatal(err)
	}
	fmt.Println(buf.String())

	//测试获取比赛
	launch.Spider_match(4)
	//测试获取配置
	val := utils2.GetVal("cookies", "Hm_lpvt_2fb6939e65e63cfbc1953f152ec2402e")
	fmt.Println(val)

	section := utils2.GetSection("cookies")
	keys := section.Keys()
	for _, e := range keys {
		fmt.Println(e.Name() + "=" + e.Value())
		fmt.Println(section.Key(e.Name()).Value())
	}
	//测试随机数
	fmt.Println(strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	//测试随机数
	for i := 0; i < 100000; i++ {
		intn := rand.Intn(10)
		if intn == 10 {
			fmt.Println("------------------------------")
			fmt.Println("------------------------------")
			fmt.Println("------------------------------")
			fmt.Println("------------------------------")
		}
		fmt.Println(intn)
	}

	//测试content长度
	fmt.Println(len("本次推荐为程序全自动处理,全程无人为参与干预.进而避免了人为分析的主观性及不稳定因素.程序根据各大波菜多维度数据,结合作者多年足球分析经验,十年程序员生涯,精雕细琢历经26个月得出的产物.程序执行流程包括且不仅限于(数据自动获取-->分析学习-->自动推送发布).经近三个月的实验准确率一直能维持在一个较高的水平.依据该项目为依托已经吸引了不少朋友,现目前通过雷速号再次验证程序的准确率,望大家长期关注校验.!"))

	//测试时间
	hours, _ := strconv.Atoi(time.Now().Format("15"))
	fmt.Println(time.Duration(int64(24 - hours)))

	//测试分析结果获取及更新

	analyService := new(service2.AnalyService)
	list := analyService.List("Euro20191206Service", 0, -1)
	result := &list[0].AnalyResult
	analyService.Modify(result)

	//测试从雷速获取可发布的比赛池
	readCloser := helper.Get(constants.MATCH_LIST_URL)
	reader := bufio.NewReader(readCloser)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break;
		} else if err != nil {
			fmt.Println(err)
			break;
		} else {
			fmt.Println(string(line))
		}
	}
	//
	poolService := new(service.MatchPoolService)
	lists := poolService.GetMatchList()
	for _, e := range lists {
		bytes, _ := json.Marshal(e)
		fmt.Println(string(bytes))
	}

}*/
