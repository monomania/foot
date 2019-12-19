package enums

/**
 *来源
 */
type PubSourceLevel int

var SourceMap = map[PubSourceLevel]string{LEISU: "雷速", WECHAT_PUBLIC: "微信公众号", SELF_MEDIA: "自媒体"}

const (
	/**
	* 雷速
	*/
	LEISU PubSourceLevel = iota
	WECHAT_PUBLIC
	SELF_MEDIA
)
