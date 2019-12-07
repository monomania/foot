package service

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/constants"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/utils"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/vo"
	"time"
)

/**
获取发布池的比赛列表
*/
type MatchPoolService struct {
	mysql.BaseService
}

/**
获取可推荐比赛列表
*/
func (this *MatchPoolService) GetMatchList() []*vo.MatchVO {
	doc, err := utils.GetDocument(constants.MATCH_LIST_URL)
	if nil != err {
		base.Log.Error("GetMatchList:" + err.Error())
		return nil
	}

	resp := make([]*vo.MatchVO, 0)
	doc.Find("tbody.data").Each(func(i int, selection *goquery.Selection) {
		matchVO := new(vo.MatchVO)
		matchVO.OddDatas = make([]vo.OddINFVO, 0)
		resp = append(resp, matchVO)
		//1.1属性获取--start
		data_id, _ := selection.Attr("data-id")
		data_sport, _ := selection.Attr("data-sport")
		data_match, _ := selection.Attr("data-match")
		data_competition, _ := selection.Attr("data-competition")
		data_zoomdate, _ := selection.Attr("data-zoomdate")
		//1.2属性获取--setin
		matchVO.DataId, _ = strconv.ParseInt(data_id, 0, 64)
		matchVO.DataSport, _ = strconv.Atoi(data_sport)
		matchVO.DataMatch, _ = strconv.ParseInt(data_match, 0, 64)
		matchVO.DataCompetition, _ = strconv.Atoi(data_competition)
		matchVO.DataZoomdate = data_zoomdate
		//1.3属性获取--end

		selection.Find("tr td").Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 0:
				//编号
				matchVO.Numb = strings.TrimSpace(selection.Text())
				break;
			case 1:
				//联赛名称
				matchVO.LeagueName = strings.TrimSpace(selection.Text())
				break;
			case 2:
				//比赛时间
				matchDate := strings.TrimSpace(selection.Text())
				match_date_stamp, _ := time.Parse("2006-01-0215:04", matchDate)
				matchVO.MatchDate = match_date_stamp
				break;
			case 3:
				//主队名称
				matchVO.MainTeam = strings.TrimSpace(selection.Text())
				break;
			case 4,5,6:
				//北单胜负过关主队赔率
				//胜平负主队赔率
				//让球胜平负主队赔率
				data_idx, _ := selection.Attr("data-idx")
				if len(data_idx) <= 0{
					break;
				}
				data_odd, _ := selection.Attr("data-odd")
				data_selects, _ := selection.Attr("data-selects")
				data_tip, _ := selection.Attr("data-tip")
				data_pk, _ := selection.Attr("data-pk")
				infvo := vo.OddINFVO{}
				infvo.DataIdx, _ = strconv.Atoi(data_idx)
				infvo.DataOdd, _ = strconv.ParseFloat(data_odd, 64)
				infvo.DataSelects, _ = strconv.Atoi(data_selects)
				infvo.DataTip = data_tip
				infvo.DataPk, _ = strconv.ParseFloat(data_pk, 64)
				matchVO.OddDatas = append(matchVO.OddDatas, infvo)
				break;
			case 7:
				//客队名称
				matchVO.GuestTeam = strings.TrimSpace(selection.Text())
				break;
			case 8,9,10:
				//北单胜负过关客队赔率
				//胜平负客队赔率
				//让球胜平负客队赔率
				data_idx, _ := selection.Attr("data-idx")
				if len(data_idx) <= 0{
					break;
				}
				data_odd, _ := selection.Attr("data-odd")
				data_selects, _ := selection.Attr("data-selects")
				data_tip, _ := selection.Attr("data-tip")
				data_pk, _ := selection.Attr("data-pk")
				infvo := vo.OddINFVO{}
				infvo.DataIdx, _ = strconv.Atoi(data_idx)
				infvo.DataOdd, _ = strconv.ParseFloat(data_odd, 64)
				infvo.DataSelects, _ = strconv.Atoi(data_selects)
				infvo.DataTip = data_tip
				infvo.DataPk, _ = strconv.ParseFloat(data_pk, 64)
				matchVO.OddDatas = append(matchVO.OddDatas, infvo)
				break;
			case 11:
				//平局
				break;
			case 12:
				break;
			case 13,14:
				//胜平负平局赔率
				//让球胜平负平局赔率
				data_idx, _ := selection.Attr("data-idx")
				if len(data_idx) <= 0{
					break;
				}
				data_odd, _ := selection.Attr("data-odd")
				data_selects, _ := selection.Attr("data-selects")
				data_tip, _ := selection.Attr("data-tip")
				data_pk, _ := selection.Attr("data-pk")
				infvo := vo.OddINFVO{}
				infvo.DataIdx, _ = strconv.Atoi(data_idx)
				infvo.DataOdd, _ = strconv.ParseFloat(data_odd, 64)
				infvo.DataSelects, _ = strconv.Atoi(data_selects)
				infvo.DataTip = data_tip
				infvo.DataPk, _ = strconv.ParseFloat(data_pk, 64)
				matchVO.OddDatas = append(matchVO.OddDatas, infvo)
			default:
				break;
			}
		})
	})
	return resp
}
