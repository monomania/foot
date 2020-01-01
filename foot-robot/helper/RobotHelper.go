package helper

import (
	"github.com/go-vgo/robotgo"
)

type RobotHelper struct {
}

/**
检查标题
 */
//func (this *RobotHelper) CheckTitle(val string) bool {
//	temp_titles := utils.GetVal("robot", "win_titiles")
//	title_arr := strings.Split(temp_titles, ",")
//
//	for _, e := range title_arr {
//		if strings.Contains(val, e) {
//			return true;
//		}
//	}
//	return false;
//}

func (this *RobotHelper) Tips() {
	robotgo.TypeStr("推荐己更新:")
	robotgo.KeyTap("enter", "control")
	robotgo.TypeStr("今日推荐,请查看https://mp.weixin.qq.com/s/yjsG81TVgWprRLKKzhon5A")
	robotgo.KeyTap("enter", "control")
	robotgo.TypeStr("赛事解析,请查看https://mp.weixin.qq.com/s/clv-Vpq-e5NxtjrryQNLhg")
	robotgo.KeyTap("enter", "control")
	robotgo.TypeStr("待选场次,请查看https://mp.weixin.qq.com/s/cVPsCefwwztkAAVW0EHZ3w")
	robotgo.KeyTap("enter", "control")
}
