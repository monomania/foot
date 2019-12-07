package main

import (
	"encoding/json"
	"fmt"
	_ "tesou.io/platform/foot-parent/foot-web/common/fliters"
	_ "tesou.io/platform/foot-parent/foot-web/common/routers"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/service"
)

func main() {
	poolService := new(service.MatchPoolService)
	list := poolService.GetMatchList()
	for _, e := range list {
		bytes, _ := json.Marshal(e)
		fmt.Println(string(bytes))
	}
	pubService := new(service.PubService)
	pubService.PubBJDC(false)

}
