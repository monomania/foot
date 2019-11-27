module tesou.io/platform/foot-parent/foot-spider

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/bitly/go-simplejson v0.5.0
	golang.org/x/net v0.0.0-20191119073136-fc4aabc6c914
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.5 // indirect
	opensource.io/go_spider v1.0.0
	tesou.io/platform/foot-parent/foot-core v1.0.0
)

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/foot-parent/foot-api => ../foot-api
	tesou.io/platform/foot-parent/foot-core => ../foot-core
)

go 1.13
