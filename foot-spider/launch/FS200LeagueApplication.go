package launch

import (
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

func Spider_league() {
	processer := proc.GetLeagueProcesser()
	processer.Startup()
}

