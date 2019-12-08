package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/service"
)

func main() {
	//测试随机数
	for i := 0; i < 10; i++ {
		intn := rand.Intn(10) + 5
		fmt.Println(intn)
	}
	//测试从雷速获取可发布的比赛池
	poolService := new(service.MatchPoolService)
	list := poolService.GetMatchList()
	for _, e := range list {
		bytes, _ := json.Marshal(e)
		fmt.Println(string(bytes))
	}

}
