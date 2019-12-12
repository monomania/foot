package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	service2 "tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/service"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/utils"
	"time"
)

func main() {
	//测试随机数
	for i := 0; i < 100000; i++ {
		intn := rand.Intn(10)
		if intn == 10{
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
	fmt.Println(time.Duration(int64(24-hours)))

	//测试分析结果获取及更新
	mysql.ShowSQL(true)
	analyService := new(service2.AnalyService)
	list := analyService.GetPubDataList("Euro81_616Service", -1)
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
		}else{
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
