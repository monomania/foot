package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	utils2 "tesou.io/platform/foot-parent/foot-core/common/utils"
	service2 "tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/service"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/utils"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"time"
)

func main() {
	//测试markdown
	tpl, err := template.ParseFiles("assets/wechat/html/week.html")
	if err != nil {
		base.Log.Error(err)
	}
	weekVO := vo.WeekVO{}
	weekVO.BeginDate = time.Now()
	weekVO.BeginDateStr = weekVO.BeginDate.Format("2006年01月02日")
	weekVO.EndDate = time.Now()
	weekVO.EndDateStr = weekVO.EndDate.Format("2006年01月02日")
	weekVO.MatchCount = 98
	weekVO.RedCount = 68
	weekVO.WalkCount = 40
	weekVO.BlackCount = 20
	weekVO.LinkRedCount = 10
	weekVO.LinkBlackCount = 5
	dataList := make([]vo.SuggestVO,1)
	suggestVO := vo.SuggestVO{}
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
	list := analyService.ListData("Euro20191206Service", 0, -1)
	result := &list[0].AnalyResult
	analyService.Modify(result)

	//测试从雷速获取可发布的比赛池
	readCloser := utils.Get(constants.MATCH_LIST_URL)
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

}
