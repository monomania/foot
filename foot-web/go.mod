module tesou.io/platform/foot-parent/foot-web

require (
	github.com/astaxie/beego v1.12.0
	github.com/lxn/walk v0.0.0-20191113135339-bf589de20b3c
	github.com/lxn/win v0.0.0-20191106123917-121afc750dd3 // indirect
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	tesou.io/platform/foot-parent/foot-api v1.0.0
	tesou.io/platform/foot-parent/foot-core v1.0.0
	tesou.io/platform/foot-parent/foot-spider v1.0.0
)

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/foot-parent/foot-api => ../foot-api
	tesou.io/platform/foot-parent/foot-core => ../foot-core
	tesou.io/platform/foot-parent/foot-spider => ../foot-spider
)

go 1.13
