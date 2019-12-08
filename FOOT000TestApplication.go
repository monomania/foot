package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/constants"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/service"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/utils"
)

func main() {
	//测试随机数
	for i := 0; i < 10; i++ {
		intn := rand.Intn(10) + 5
		fmt.Println(intn)
	}
	//测试从雷速获取可发布的比赛池
	readCloser := utils.Get(constants.MATCH_LIST_URL)
	reader := bufio.NewReader(readCloser)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break;
		} else if err != nil {
			fmt.Println(err)
		}else{
			fmt.Println(string(line))
		}
	}
	//
	poolService := new(service.MatchPoolService)
	list := poolService.GetMatchList()
	for _, e := range list {
		bytes, _ := json.Marshal(e)
		fmt.Println(string(bytes))
	}

}
