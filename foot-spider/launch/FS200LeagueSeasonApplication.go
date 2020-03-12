package launch

import (
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

func Spider_leagueSeason() {
	processer := proc.GetLeagueSeasonProcesser()
	processer.Startup()
}

