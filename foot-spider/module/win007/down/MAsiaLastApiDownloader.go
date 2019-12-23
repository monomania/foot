package down

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	//    iconv "github.com/djimenez/iconv-go"
	"github.com/hu17889/go_spider/core/common/mlog"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/common/util"
	//    "golang.org/x/text/encoding/simplifiedchinese"
	//    "golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	//"fmt"
	"golang.org/x/net/html/charset"
	//    "regexp"
	//    "golang.org/x/net/html"
	"compress/gzip"
)

// The MAsiaLastApiDownloader download page by package net/http.
// The "html" content is contained in dom parser of package goquery.
// The "json" content is saved.
// The "jsonp" content is modified to json.
// The "text" content will save body plain text only.
// The page result is saved in Page.
type MAsiaLastApiDownloader struct {
}

func NewMAsiaLastApiDownloader() *MAsiaLastApiDownloader {
	return &MAsiaLastApiDownloader{}
}

func (this *MAsiaLastApiDownloader) Download(req *request.Request) *page.Page {
	var mtype string
	var p = page.NewPage(req)
	mtype = req.GetResponceType()
	switch mtype {
	case "html":
		return this.downloadHtml(p, req)
	case "json":
		fallthrough
	case "jsonp":
		return this.downloadJson(p, req)
	case "text":
		return this.downloadText(p, req)
	default:
		mlog.LogInst().LogError("error request type:" + mtype)
	}
	return p
}

// Charset auto determine. Use golang.org/x/net/html/charset. Get page body and change it to utf-8
func (this *MAsiaLastApiDownloader) changeCharsetEncodingAuto(contentTypeStr string, sor io.ReadCloser) string {
	var err error
	destReader, err := charset.NewReader(sor, contentTypeStr)

	if err != nil {
		mlog.LogInst().LogError(err.Error())
		destReader = sor
	}

	var sorbody []byte
	if sorbody, err = ioutil.ReadAll(destReader); err != nil {
		mlog.LogInst().LogError(err.Error())
		// For gb2312, an error will be returned.
		// Error like: simplifiedchinese: invalid GBK encoding
		// return ""
	}
	//e,name,certain := charset.DetermineEncoding(sorbody,contentTypeStr)
	bodystr := string(sorbody)

	return bodystr
}

func (this *MAsiaLastApiDownloader) changeCharsetEncodingAutoGzipSupport(contentTypeStr string, sor io.ReadCloser) string {
	var err error
	gzipReader, err := gzip.NewReader(sor)
	if err != nil {
		mlog.LogInst().LogError(err.Error())
		return ""
	}
	defer gzipReader.Close()
	destReader, err := charset.NewReader(gzipReader, contentTypeStr)

	if err != nil {
		mlog.LogInst().LogError(err.Error())
		destReader = sor
	}

	var sorbody []byte
	if sorbody, err = ioutil.ReadAll(destReader); err != nil {
		mlog.LogInst().LogError(err.Error())
		// For gb2312, an error will be returned.
		// Error like: simplifiedchinese: invalid GBK encoding
		// return ""
	}
	//e,name,certain := charset.DetermineEncoding(sorbody,contentTypeStr)
	bodystr := string(sorbody)

	return bodystr
}

func (this *MAsiaLastApiDownloader) getHeader() http.Header {
	//设置head
	header := http.Header{}
	header.Add("Host", "m.win007.com")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	header.Add("Accept", "*/*")
	header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	header.Add("Accept-Encoding", "gzip, deflate")
	header.Add("Connection", "keep-alive")
	header.Add("Referer", "http://m.win007.com/")
	header.Add("Cookie", "UM_distinctid=16d41e1fceb179-0546f41e76d7cd-14377840-1aeaa0-16d41e1fcec16b; CNZZDATA1643862=cnzz_eid%3D1697126684-1568763936-%26ntime%3D1573196700")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	return header
}

// Download file and change the charset of page charset.
func (this *MAsiaLastApiDownloader) downloadFile(p *page.Page, req *request.Request) (*page.Page, string) {
	var err error
	var urlstr string

	req.Header = this.getHeader()
	if urlstr = req.GetUrl(); len(urlstr) == 0 {
		mlog.LogInst().LogError("url is empty")
		p.SetStatus(true, "url is empty")
		return p, ""
	}

	var resp *http.Response

	if proxystr := req.GetProxyHost(); len(proxystr) != 0 {
		//using http proxy
		//fmt.Print("HttpProxy Enter ",proxystr,"\n")
		resp, err = connectByHttpProxy(p, req)
	} else {
		//normal http download
		//fmt.Print("Http Normal Enter \n",proxystr,"\n")
		resp, err = connectByHttp(p, req)
	}

	if err != nil {
		return p, ""
	}

	//b, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("Resp body %v \r\n", string(b))

	p.SetHeader(resp.Header)
	p.SetCookies(resp.Cookies())

	// get converter to utf-8
	var bodyStr string
	if resp.Header.Get("Content-Encoding") == "gzip" {
		bodyStr = this.changeCharsetEncodingAutoGzipSupport(resp.Header.Get("Content-Type"), resp.Body)
	} else {
		bodyStr = this.changeCharsetEncodingAuto(resp.Header.Get("Content-Type"), resp.Body)
	}
	//fmt.Printf("utf-8 body %v \r\n", bodyStr)
	defer resp.Body.Close()
	return p, bodyStr
}

func (this *MAsiaLastApiDownloader) downloadHtml(p *page.Page, req *request.Request) *page.Page {
	var err error
	p, destbody := this.downloadFile(p, req)
	//fmt.Printf("Destbody %v \r\n", destbody)
	if !p.IsSucc() {
		//fmt.Print("Page error \r\n")
		return p
	}
	bodyReader := bytes.NewReader([]byte(destbody))

	var doc *goquery.Document
	if doc, err = goquery.NewDocumentFromReader(bodyReader); err != nil {
		mlog.LogInst().LogError(err.Error())
		p.SetStatus(true, err.Error())
		return p
	}

	var body string
	if body, err = doc.Html(); err != nil {
		mlog.LogInst().LogError(err.Error())
		p.SetStatus(true, err.Error())
		return p
	}

	p.SetBodyStr(body).SetHtmlParser(doc).SetStatus(false, "")

	return p
}

func (this *MAsiaLastApiDownloader) downloadJson(p *page.Page, req *request.Request) *page.Page {
	var err error
	p, destbody := this.downloadFile(p, req)
	if !p.IsSucc() {
		return p
	}

	var body []byte
	body = []byte(destbody)
	mtype := req.GetResponceType()
	if mtype == "jsonp" {
		tmpstr := util.JsonpToJson(destbody)
		body = []byte(tmpstr)
	}

	var r *simplejson.Json
	if r, err = simplejson.NewJson(body); err != nil {
		mlog.LogInst().LogError(string(body) + "\t" + err.Error())
		p.SetStatus(true, err.Error())
		return p
	}

	// json result
	p.SetBodyStr(string(body)).SetJson(r).SetStatus(false, "")

	return p
}

func (this *MAsiaLastApiDownloader) downloadText(p *page.Page, req *request.Request) *page.Page {
	p, destbody := this.downloadFile(p, req)
	if !p.IsSucc() {
		return p
	}

	p.SetBodyStr(destbody).SetStatus(false, "")
	return p
}
