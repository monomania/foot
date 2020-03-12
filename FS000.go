package main

import (
	"fmt"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"time"
)

func main() {
	now := time.Now()
	temp_year, _ := time.ParseDuration("-8760h")
	add := now.Add(temp_year)
	parse:= now.Format( "2006")
	parse2 := add.Format( "2006")
	fmt.Println(parse)
	fmt.Println(parse2)

	//launch.Spider_league()
	launch.Spider_leagueSeason()
}

