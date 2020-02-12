package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	service4 "tesou.io/platform/foot-parent/foot-core/module/analy/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/core/service"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/constants"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/service"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/utils"
	service3 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-core/module/wechat/controller"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"time"
)

func init() {

}

func main() {
	var input string
	if len(os.Args) > 1 {
		input = strings.ToLower(os.Args[1])
	} else {
		input = ""
	}

	switch input {
	case "exit\n", "exit", "quit\n", "quit":
		break;
	case "\n", "":
	case "init\n", "init":
		launch2.GenTable()
		launch2.TruncateTable()
	case "spider\n", "spider":
		launch.Spider()
	case "analy\n", "analy":
		launch2.Analy()
	case "mr\n", "mr":
		//更新结果
		analyService := new(service4.AnalyService)
		analyService.ModifyAllResult()
	case "baseFace\n", "baseFace":
		launch.Spider_match(4)
		launch.Spider_asiaLastNew(true)
		launch.Spider_baseFace(true)
	case "limit\n", "limit":
		pubLimitService := new(service.PubLimitService)
		publimit := pubLimitService.GetPublimit()
		bytes, _ := json.Marshal(publimit)
		fmt.Println("发布限制信息为:" + string(bytes))
	case "price\n", "price":
		priceService := new(service.PriceService)
		price := priceService.GetPrice()
		bytes, _ := json.Marshal(price)
		fmt.Println("收费信息为:" + string(bytes))
	case "matchpool\n", "matchpool":
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
		//尝试获取比赛列表
		poolService := new(service.MatchPoolService)
		list := poolService.GetMatchList()
		for _, e := range list {
			bytes, _ := json.Marshal(e)
			fmt.Println(string(bytes))
		}
	case "pub\n", "pub":
		pubService := new(service.PubService)
		pubService.PubBJDC()
	case "autospider\n", "autospider":
		for {
			base.Log.Info("--------数据更新开始运行--------")
			//1.安装数据库
			//2.配置好数据库连接,打包程序发布
			//3.程序执行流程,周期定制定为一天三次
			//3.1 FS000Application 爬取数据
			launch.Spider()
			//3.2 FC002AnalyApplication 分析得出推荐列表
			launch2.Analy()
			configService := new(service2.ConfService)
			base.Log.Info("--------数据更新周期结束--------")
			time.Sleep(time.Duration(configService.GetSpiderCycleTime()) * time.Minute)
		}
	case "autoleisu\n", "autoleisu":
		for {
			base.Log.Info("--------发布开始运行--------")
			//3.3 FW001PubApplication 执行发布到雷速
			pubService := new(service.PubService)
			pubService.PubBJDC()
			base.Log.Info("--------发布周期结束--------")
			time.Sleep(time.Duration(pubService.CycleTime()) * time.Minute)
		}
	case "autowechat\n", "autowechat":
		for {
			base.Log.Info("--------发布公众号开始运行--------")
			//3.3 FW001PubApplication 执行发布到雷速
			materialController := new(controller.MaterialController)
			materialController.ModifyNewsOnly()
			base.Log.Info("--------发布公众号周期结束--------")
			time.Sleep(10 * time.Minute)
		}
	case "auto\n", "auto":
		go func() {
			for {
				base.Log.Info("--------全量数据更新开始运行--------")
				//1.安装数据库
				//2.配置好数据库连接,打包程序发布
				//3.程序执行流程,周期定制定为一天三次
				//3.1 FS000Application 爬取数据
				//清空数据库数据,为爬虫作准备
				launch.Clean()
				launch.Spider()
				//3.2 FC002AnalyApplication 分析得出推荐列表
				launch2.Analy()
				configService := new(service2.ConfService)
				base.Log.Info("--------全量数据更新周期结束--------")

				base.Log.Info("--------全量比赛发布公众号开始运行--------")
				//3.3 FW001PubApplication 执行发布到雷速
				materialController := new(controller.MaterialController)
				materialController.ModifyNewsOnly()
				base.Log.Info("--------全量临场比赛发布公众号周期结束--------")

				time.Sleep(time.Duration(configService.GetSpiderCycleTime()) * time.Minute)
			}
		}()
		//go func() {
		//	for {
		//		base.Log.Info("--------发布开始运行--------")
		//		//3.3 FW001PubApplication 执行发布到雷速
		//		pubService := new(service.PubService)
		//		pubService.PubBJDC()
		//		base.Log.Info("--------发布周期结束--------")
		//		time.Sleep(time.Duration(pubService.CycleTime()) * time.Minute)
		//	}
		//}()
		time.Sleep(1 * time.Second)
		for {
			matchLastService := new(service3.MatchLastService)
			matchLasts := matchLastService.FindNear()
			if len(matchLasts) > 0 {
				base.Log.Info("--------临场比赛,共", len(matchLasts), "场,更新开始运行--------")
				launch.Spider_Near()
				launch2.Analy_Near()
				base.Log.Info("--------临场比赛,共", len(matchLasts), "场,更新周期结束--------")

				base.Log.Info("--------临场比赛发布公众号开始运行--------")
				//3.3 FW001PubApplication 执行发布到雷速
				materialController := new(controller.MaterialController)
				materialController.ModifyNewsOnly()
				base.Log.Info("--------临场比赛发布公众号周期结束--------")
			}
			time.Sleep(5 * time.Minute)
		}

	default:

	}

}
