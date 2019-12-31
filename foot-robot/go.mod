module tesou.io/platform/foot-parent/foot-robot

require github.com/go-vgo/robotgo v0.0.0-20191226160149-28f256a4c5a0

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/foot-parent/foot-api => ../foot-api
	tesou.io/platform/foot-parent/foot-core => ../foot-core
	tesou.io/platform/foot-parent/foot-spider => ../foot-spider
)

go 1.13
