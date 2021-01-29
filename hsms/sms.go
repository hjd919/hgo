package gom

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func NewSms(conf *SmsConf) *Sms {
	return &Sms{
		conf: conf,
	}
}

type SmsConf struct {
	AccessKeyId     string // id
	AccessKeySecret string // secret
	SignName        string // 短信签名
	TemplateCode    string // 模版code
}

type Sms struct {
	conf *SmsConf
}

// 发送短信
type SmsSendParam struct {
	Phones        []string               //发送的手机号 18500223333
	TemplateParam map[string]interface{} //模版替换的内容 {code:1234}
}
type SmsReply struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

func (t *Sms) Send(p *SmsSendParam) error {
	templateParam := JsonEncode(p.TemplateParam)

	//组织参数
	params := map[string]string{
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureNonce":   fmt.Sprintf("%d", rand.Int63()),
		"AccessKeyId":      t.conf.AccessKeyId,
		"SignatureVersion": "1.0",
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           "JSON",
		"Action":           "SendSms",
		"Version":          "2017-05-25",
		"RegionId":         "cn-hangzhou",
		"PhoneNumbers":     strings.Join(p.Phones, ","), //多个手机号，","相隔
		"SignName":         t.conf.SignName,
		"TemplateParam":    templateParam,
		"TemplateCode":     t.conf.TemplateCode,
	}
	//排序key
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//组织字符串
	var sortQueryString string
	for _, v := range keys {
		sortQueryString = fmt.Sprintf("%s&%s=%s", sortQueryString, t.replace(v), t.replace(params[v]))
	}
	//组织sign
	stringToSign := fmt.Sprintf("GET&%s&%s", t.replace("/"), t.replace(sortQueryString[1:]))
	mac := hmac.New(sha1.New, []byte(fmt.Sprintf("%s&", t.conf.AccessKeySecret)))
	mac.Write([]byte(stringToSign))
	sign := t.replace(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	//发送访问请求
	strUrl := fmt.Sprintf("http://dysmsapi.aliyuncs.com/?Signature=%s%s", sign, sortQueryString)
	resp, err := http.Get(strUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//组织返回数据
	ssr := &SmsReply{}
	if err := json.Unmarshal(body, ssr); err != nil {
		return err
	}
	if ssr.Code == "SignatureNonceUsed" {
		return t.Send(p)
	} else if ssr.Code != "OK" {
		return errors.New(ssr.Code + ":" + ssr.Message)
	}
	return nil
}

//替换字符串
func (t *Sms) replace(in string) string {
	rep := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	return rep.Replace(url.QueryEscape(in))
}
