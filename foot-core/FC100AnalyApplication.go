package main

import (
	"log"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

func main() {
	Analy()
}

func Analy() []interface{} {
	analysisService := new(service.AnalyService)
	analysisService.MaxLetBall = 0.75
	analysisService.PrintOddData = false
	log.Println("-----------------------------------------------")
	log.Println("----------------计算欧86之差-------------------")
	log.Println("-----------------------------------------------")
	r1 := analysisService.Euro_Calc()
	log.Println("-----------------------------------------------")
	log.Println("---------------计算亚欧之差--------------------")
	log.Println("-----------------------------------------------")
	r2 := analysisService.Euro_Asia_Diff()

	i := append(r1, r2)
	return i

}
