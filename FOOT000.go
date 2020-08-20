package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
	"sort"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	service2 "tesou.io/platform/foot-parent/foot-core/module/core/service"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"time"
)

func init() {

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configService := new(service2.ConfService)

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang, l",
				Value: "english",
				Usage: "Language for the greeting",
			},
			//&cli.StringFlag{
			//	Name:  "config, c",
			//	Usage: "Load configuration from `FILE`",
			//},
		},
		Commands: []*cli.Command{
			{
				Name:    "initdb",
				Aliases: []string{"i", "db"},
				Usage:   "初始化数据库表",
				Action: func(c *cli.Context) error {
					launch2.GenTable()
					//launch2.TruncateTable()
					return nil
				},
			},
			{
				Name:    "spider",
				Aliases: []string{"s", "sp"},
				Usage:   "抓取数据",
				Subcommands: []*cli.Command{
					{
						Name:  "default",
						Usage: "default",
						Action: func(c *cli.Context) error {
							launch.Spider()
							return nil
						},
					},
				},
			},
			{
				Name:    "analy",
				Usage:   "分析数据",
				Subcommands: []*cli.Command{
					{
						Name:  "curent",
						Usage: "分析当前数据",
						Action: func(c *cli.Context) error {
							launch2.Analy(false)
							return nil
						},
					},
					{
						Name:  "all",
						Usage: "分析所有数据",
						Action: func(c *cli.Context) error {
							launch2.Analy(true)
							return nil
						},
					},
					{
						Name:  "auto",
						Usage: "定时进行分析当前",
						Action: func(c *cli.Context) error {
							for {
								launch2.Analy(false)
								time.Sleep(time.Duration(configService.GetSpiderCycleTime()) * time.Minute)
							}
							return nil
						},
					},
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
