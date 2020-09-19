module tesou.io/platform/foot-parent

require (
	github.com/astaxie/beego v1.12.0
	github.com/go-vgo/robotgo v0.0.0-20191226160149-28f256a4c5a0
	github.com/lxn/walk v0.0.0-20191128110447-55ccb3a9f5c1 // indirect
	github.com/urfave/cli/v2 v2.2.0
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	tesou.io/platform/foot-parent/foot-api v1.0.0
	tesou.io/platform/foot-parent/foot-core v1.0.0
	tesou.io/platform/foot-parent/foot-spider v1.0.0
)

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/foot-parent/foot-api => ./foot-api
	tesou.io/platform/foot-parent/foot-core => ./foot-core
	tesou.io/platform/foot-parent/foot-robot => ./foot-robot
	tesou.io/platform/foot-parent/foot-spider => ./foot-spider
	tesou.io/platform/foot-parent/foot-web => ./foot-web
)

go 1.13
