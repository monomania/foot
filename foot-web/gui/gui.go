package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"tesou.io/platform/foot-parent/foot-core/launch"
	launch2 "tesou.io/platform/foot-parent/foot-spider/launch"
)

func main() {
	build_win()
}

func build_win() {
	var inTE *walk.TextEdit
	window := MainWindow{
		Title:   "Foot",
		MinSize: Size{400, 300},
		Layout:  VBox{},
		Children: []Widget{
			TextEdit{AssignTo: &inTE, ReadOnly: true, HScroll: true},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Clean数据",
						OnClicked: func() {
							inTE.SetText("同步清理数据...")
							go launch.DBInit()
						},
					},
					PushButton{
						Text: "Spider数据",
						OnClicked: func() {
							inTE.SetText("抓取数据...")
							go launch2.Spider(4)
						},
					},
					PushButton{
						Text: "Analy数据",
						OnClicked: func() {
							inTE.SetText("分析数据...")
							go launch2.Spider(4)
						},
					},
				},
			},
		},
	}
	window.Run()
}
