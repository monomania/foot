package utils

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
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
	setGetHeader(request)
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

func Struct2Map(obj interface{}) map[string]string {
	obj_v := reflect.ValueOf(obj)
	v := obj_v.Elem()
	typeOfType := v.Type()
	var data = make(map[string]string)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		marshal, _ := json.Marshal(field.Interface())
		data[strings.ToLower(typeOfType.Field(i).Name)] = string(marshal)
	}
	return data
}

/**
 *
 */
func Post(apiUrl string, data interface{}) string {
	struct2Map := Struct2Map(data)
	values := url.Values{}
	for k, v := range struct2Map {
		values.Add(k,v)
	}
	encode := values.Encode()
	fmt.Println(encode)

	client := &http.Client{}
	u, _ := url.ParseRequestURI(apiUrl)
	request, err := http.NewRequest("POST", u.String(), strings.NewReader(encode))
	if nil != err {
		base.Log.Error(err)
	}

	//设置请求头
	setPostHeader(request)
	request.Header.Add("Content-Length", strconv.Itoa(len(encode)))
	response, err := client.Do(request)
	if nil != err {
		base.Log.Error(err)
	}
	var reader io.ReadCloser
	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			base.Log.Error("PubPost:" + err.Error())
			return ""
		}
	} else {
		reader = response.Body
	}
	bytes, e := ioutil.ReadAll(reader)
	if e != nil {
		base.Log.Error("PubPost:" + e.Error())
		return ""
	}
	return string(bytes)
}

func setGetHeader(req *http.Request) {
	//设置cookies
	setCookies(req)
	//设置head
	req.Header.Add("Host", "hao.leisu.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://hao.leisu.com/match")
	//req.Header.Add("Cookie", "Hm_lvt_63b82ac6d9948bad5e14b1398610939a=1574214476,1574241706,1574912834; LWT=J38HWCiN9ar%2B5Ih2MPMpuOsXNUBwaQIkBg5DEfwTblTtgQFlCxIKOl8yHndiSdTBG2d3G%2BMfhbxNz1pPdnuxQMwxhPjiHrwoqUOAycHWgMA%3D; Hm_lvt_2fb6939e65e63cfbc1953f152ec2402e=1574238486,1574241710,1574241711,1574912837; Hm_lpvt_63b82ac6d9948bad5e14b1398610939a=1574912834; acw_tc=2f61f27615749128339236126e4d79296e8377930295ed12ea7a883b6b8e6f; SERVERID=4ab2f7c19b72630dd03ede01228e3e61|1574914772|1574912833; Hm_lpvt_2fb6939e65e63cfbc1953f152ec2402e=1574914777")
	req.Header.Add("Cookie", "Hm_lvt_63b82ac6d9948bad5e14b1398610939a=1574241706,1574912834,1575529351,1575748957; Hm_lvt_2fb6939e65e63cfbc1953f152ec2402e=1574241711,1574912837,1575529353,1575748959; acw_tc=2f61f27615749128339236126e4d79296e8377930295ed12ea7a883b6b8e6f; Hm_lpvt_63b82ac6d9948bad5e14b1398610939a=1575748957; SERVERID=781c670734bb86bae9c144a7d6e78b87|1575805618|1575805541; Hm_lpvt_2fb6939e65e63cfbc1953f152ec2402e=1575805623; acw_sc__v2=5dece2524bac33b7e669c181b3e0c72ebd65e176; LWT=RIW0BQy19w%2FnuztwzW1REnAawhWx06ZglggjUPnfZl2Pf8FnjDk%2Fnx8Bbfa%2Fk4IhxfP69hKNjHuiSCXtYO8rCaHgDYEpu%2FItWxCIpZpz67o%3D")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
}

func setPostHeader(req *http.Request) {
	//设置cookies
	setCookies(req)
	//设置head
	req.Header.Add("Host", "api.leisu.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Add("Content-Length", "1624")
	req.Header.Add("Origin", "https://hao.leisu.com")
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://hao.leisu.com/match")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
}


func setCookies(req *http.Request){
	req.AddCookie(&http.Cookie{Name:"acw_tc",Value:"2760776115742144718355201e7cb9413aba523148a7028556635a156dd569"})
	req.AddCookie(&http.Cookie{Name:"Hm_lpvt_63b82ac6d9948bad5e14b1398610939a",Value:"1575748957"})
	req.AddCookie(&http.Cookie{Name:"Hm_lvt_63b82ac6d9948bad5e14b1398610939a",Value:"1574241706,1574912834,1575529351,1575748957"})
	req.AddCookie(&http.Cookie{Name:"LWT",Value:"RIW0BQy19w/nuztwzW1REnAawhWx06ZglggjUPnfZl2Pf8FnjDk/nx8Bbfa/k4IhxfP69hKNjHuiSCXtYO8rCaHgDYEpu/ItWxCIpZpz67o="})
	req.AddCookie(&http.Cookie{Name:"SERVERID",Value:"4ab2f7c19b72630dd03ede01228e3e61|1575815203|1575815178"})
}