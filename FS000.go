package main

import (
	"fmt"
	"time"
)

var SingleThread bool

func main() {
	now := time.Now()
	temp_year, _ := time.ParseDuration("-8760h")
	add := now.Add(temp_year)
	parse := now.Format("2006")
	parse2 := add.Format("2006")
	fmt.Println(parse)
	fmt.Println(parse2)
	//launch.Spider_match_his("2019")
	//launch.Spider_History()
	//launch.Spider_leagueSeason()

	fmt.Println(SingleThread)

}
