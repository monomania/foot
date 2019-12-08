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
	//req.Header.Add("Cookie", "Hm_lvt_63b82ac6d9948bad5e14b1398610939a=1574284240,1575125973; acw_tc=2760774915751259641277888e0c0d3b7dc1faefe284f97f4e153b9ff71c51; LWT=hyYXzEENtV83OvKqggGZwgvmX1ld25H7RJDz92A1QcddMsndkOTK2Q7F4cbUp3M2XEzCr06PzaAzQeuyk%2B93RUcI5naHo4rL0ArpE%2B%2F65eU%3D; SERVERID=b1339a6cb30fad3b30cae2f79c06f0ea|1575726984|1575726466")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
}


func setCookies(req *http.Request){
	req.AddCookie(&http.Cookie{Name:"acw_sc__v2",Value:"5ded0fc3c5dcf635e4781179385f6f75784750b7"})
	req.AddCookie(&http.Cookie{Name:"acw_tc",Value:"2f61f27615749128339236126e4d79296e8377930295ed12ea7a883b6b8e6f"})
	req.AddCookie(&http.Cookie{Name:"Hm_lpvt_2fb6939e65e63cfbc1953f152ec2402e",Value:"1575815243"})
	req.AddCookie(&http.Cookie{Name:"Hm_lpvt_63b82ac6d9948bad5e14b1398610939a",Value:"1575748957"})
	req.AddCookie(&http.Cookie{Name:"Hm_lvt_2fb6939e65e63cfbc1953f152ec2402e",Value:"1574241711,1574912837,1575529353,1575748959"})
	req.AddCookie(&http.Cookie{Name:"Hm_lvt_63b82ac6d9948bad5e14b1398610939a",Value:"1574241706,1574912834,1575529351,1575748957"})
	req.AddCookie(&http.Cookie{Name:"LWT",Value:"KBjo3NZHBCcMfcnaZ1JXVula6ZKdmthk0DqnddpGOcwxb/d/nr2ULipCnjSsetvzCgwBKwkj7nBLrOAdTnObJULe9q6AAzwXSGYnfRLNE2U="})
	req.AddCookie(&http.Cookie{Name:"SERVERID",Value:"f35c9b1c268fcc8b58043d36d0dddd26|1575817686|1575815177"})
}