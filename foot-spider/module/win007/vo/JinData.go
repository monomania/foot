package vo

import "tesou.io/platform/foot-parent/foot-api/module/match/pojo"

type JinData struct {
	HomeInfo []*pojo.BFJin `json:"HomeInfo"`
	GuestInfo []*pojo.BFJin `json:"GuestInfo"`
}
