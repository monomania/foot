module tesou.io/platform/foot-parent/foot-gui

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/foot-parent/foot-api => ../foot-api
	tesou.io/platform/foot-parent/foot-core => ../foot-core
	tesou.io/platform/foot-parent/foot-spider => ../foot-spider
)

go 1.13
