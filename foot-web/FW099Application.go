package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"

	"net/http"
	_ "tesou.io/platform/foot-parent/foot-web/common/fliters"
	_ "tesou.io/platform/foot-parent/foot-web/common/routers"
	"time"
)

func main() {
	MatchInfo()
}

//登录模拟
func Login() {

}

//获取可发布单关的比赛信息
func MatchInfo() {
	url := "https://hao.leisu.com/match"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if nil != err {
		base.Log.Error(err)
	}
	setHeader(request)
	response, err := client.Do(request)
	if nil != err {
		base.Log.Error(err)
	}
	var reader io.ReadCloser
	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return
		}
	} else {
		reader = response.Body
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Text())
	fmt.Println(doc.Html())

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

//

func Post(url string, data interface{}, contentType string) string {

	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(`POST`, url, bytes.NewBuffer(jsonStr))
	req.Header.Add(`content-type`, contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func setHeader(req *http.Request) {
	//设置head
	req.Header.Add("Host","hao.leisu.com")
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language","zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding","gzip, deflate, br")
	req.Header.Add("Connection","keep-alive")
	req.Header.Add("Referer","https://hao.leisu.com/match")
	req.Header.Add("Cookie","Hm_lvt_63b82ac6d9948bad5e14b1398610939a=1574214476,1574241706,1574912834; LWT=J38HWCiN9ar%2B5Ih2MPMpuOsXNUBwaQIkBg5DEfwTblTtgQFlCxIKOl8yHndiSdTBG2d3G%2BMfhbxNz1pPdnuxQMwxhPjiHrwoqUOAycHWgMA%3D; Hm_lvt_2fb6939e65e63cfbc1953f152ec2402e=1574238486,1574241710,1574241711,1574912837; Hm_lpvt_63b82ac6d9948bad5e14b1398610939a=1574912834; acw_tc=2f61f27615749128339236126e4d79296e8377930295ed12ea7a883b6b8e6f; SERVERID=4ab2f7c19b72630dd03ede01228e3e61|1574914772|1574912833; Hm_lpvt_2fb6939e65e63cfbc1953f152ec2402e=1574914777")
	req.Header.Add("Upgrade-Insecure-Requests","1")
	req.Header.Add("Pragma","no-cache")
	req.Header.Add("Cache-Control","no-cache")
}

func getCookies() []http.Cookie{
	cookies := make([]http.Cookie,5)
	cookie1 := http.Cookie{}
	cookie1.Name ="acw_tc"
	cookie1.Value ="2f61f27615749128339236126e4d79296e8377930295ed12ea7a883b6b8e6f"
	cookie2 := http.Cookie{}
	cookie2.Name ="Hm_lpvt_2fb6939e65e63cfbc1953f152ec2402e"
	cookie2.Value ="1574914720"
	cookie3 := http.Cookie{}
	cookie3.Name ="Hm_lpvt_63b82ac6d9948bad5e14b1398610939a"
	cookie3.Value ="1574912834"
	cookie4 := http.Cookie{}
	cookie4.Name ="Hm_lvt_2fb6939e65e63cfbc1953f152ec2402e"
	cookie4.Value ="1574238486,1574241710,1574241711,1574912837"
	cookie5 := http.Cookie{}
	cookie5.Name ="Hm_lvt_63b82ac6d9948bad5e14b1398610939a"
	cookie5.Value ="1574214476,1574241706,1574912834"
	cookie6 := http.Cookie{}
	cookie6.Name ="LWT"
	cookie6.Value ="J38HWCiN9ar+5Ih2MPMpuOsXNUBwaQIkBg5DEfwTblTtgQFlCxIKOl8yHndiSdTBG2d3G+MfhbxNz1pPdnuxQMwxhPjiHrwoqUOAycHWgMA="
	cookie7 := http.Cookie{}
	cookie7.Name ="SERVERID"
	cookie7.Value ="4ab2f7c19b72630dd03ede01228e3e61|1574914716|1574912833"

	cookies = append(cookies,cookie1)
	cookies = append(cookies,cookie2)
	cookies = append(cookies,cookie3)
	cookies = append(cookies,cookie4)
	cookies = append(cookies,cookie5)
	return cookies
}