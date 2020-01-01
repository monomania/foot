package utils

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
)

/**
检查标题
 */
func CheckTitle(val string) bool{
	temp_titles := utils.GetVal("robot", "win_titiles")
	title_arr := strings.Split(temp_titles, ",")

	for _, e := range title_arr {
		if strings.Contains(val,e){
			return true;
		}
	}
	return false;
}



