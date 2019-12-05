package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"tesou.io/platform/foot-parent/foot-api/common/base"
)

/**
 *
 */
func Get(url string) io.ReadCloser {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if nil != err {
		base.Log.Error(err)
	}

	//设置请求头
	setHeader(request)
	response, err := client.Do(request)
	if nil != err {
		base.Log.Error(err)
	}
	var reader io.ReadCloser
	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			base.Log.Error("GetDocument:" + err.Error())
			return nil
		}
	} else {
		reader = response.Body
	}
	return reader
}
func GetText(url string) string {
	reader := Get(url)
	bytes, e := ioutil.ReadAll(reader)
	if e != nil {
		base.Log.Error("GetText:" + e.Error())
		return ""
	}
	return string(bytes)
}
func GetDocument(url string) (*goquery.Document, error) {
	reader := Get(url)
	return goquery.NewDocumentFromReader(reader)
}


/**
 *
 */
func Post(url string, data interface{}) string {
	byteData, _ := json.Marshal(data)
	params := bytes.NewReader(byteData)

	client := &http.Client{}
	request, err := http.NewRequest("Post", url, params)
	if nil != err {
		base.Log.Error(err)
	}

	//设置请求头
	setHeader(request)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if nil != err {
		base.Log.Error(err)
	}
	var reader io.ReadCloser
	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			base.Log.Error("Post:" + err.Error())
			return ""
		}
	} else {
		reader = response.Body
	}
	bytes, e := ioutil.ReadAll(reader)
	if e != nil {
		base.Log.Error("Post:" + e.Error())
		return ""
	}
	return string(bytes)
}

func setHeader(req *http.Request) {
	//设置head
	req.Header.Add("Host", "hao.leisu.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://hao.leisu.com/match")
	req.Header.Add("Cookie", "Hm_lvt_63b82ac6d9948bad5e14b1398610939a=1574214476,1574241706,1574912834; LWT=J38HWCiN9ar%2B5Ih2MPMpuOsXNUBwaQIkBg5DEfwTblTtgQFlCxIKOl8yHndiSdTBG2d3G%2BMfhbxNz1pPdnuxQMwxhPjiHrwoqUOAycHWgMA%3D; Hm_lvt_2fb6939e65e63cfbc1953f152ec2402e=1574238486,1574241710,1574241711,1574912837; Hm_lpvt_63b82ac6d9948bad5e14b1398610939a=1574912834; acw_tc=2f61f27615749128339236126e4d79296e8377930295ed12ea7a883b6b8e6f; SERVERID=4ab2f7c19b72630dd03ede01228e3e61|1574914772|1574912833; Hm_lpvt_2fb6939e65e63cfbc1953f152ec2402e=1574914777")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
}
