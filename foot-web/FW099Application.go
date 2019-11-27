package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	_ "tesou.io/platform/foot-parent/foot-web/common/fliters"
	_ "tesou.io/platform/foot-parent/foot-web/common/routers"
	"time"
)

func main() {

	//url := "https://api.leisu.com/api/v2/user/login"
	//contentType := "application/json, text/javascript, */*; q=0.01"


}

// 发送POST请求
// url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
// content:请求放回的内容
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
