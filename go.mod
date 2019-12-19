module tesou.io/platform/foot-parent

require (
	github.com/astaxie/beego v1.12.0
	tesou.io/platform/foot-parent/foot-api v1.0.0
	tesou.io/platform/foot-parent/foot-core v1.0.0
	tesou.io/platform/foot-parent/foot-spider v1.0.0
)

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/foot-parent/foot-api => ./foot-api
	tesou.io/platform/foot-parent/foot-core => ./foot-core
	tesou.io/platform/foot-parent/foot-spider => ./foot-spider
	tesou.io/platform/foot-parent/foot-web => ./foot-web
)

go 1.13
