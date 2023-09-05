package cat

import (
	"github.com/ciyjbo/buildHttps"
	"github.com/tidwall/gjson"
	"log"
)

type Ciweimao struct {
	params    map[string]string
	host      string
	version   string
	decodeKey string
}

func InitCiweimaoClient(account, token string) *Ciweimao {
	if len(token) != 32 {
		log.Fatal("token格式错误")
	}
	return &Ciweimao{
		host:      "https://app.hbooker.com",
		version:   "2.9.290",
		decodeKey: "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn",
		params: map[string]string{
			"device_token": "ciweimao_",
			"app_version":  "2.9.290",
			"login_token":  token,
			"account":      account,
		},
	}
}

func (cat *Ciweimao)NewCiweimaoParams(account, token string){
	cat.params["account"] = account
	cat.params["login_token"] = token
}

func (cat *Ciweimao) post(url string, data map[string]string) gjson.Result {
	if data != nil {
		for k, v := range data {
			cat.params[k] = v
		}
	}
	response := buildHttps.Get(cat.host+url, cat.params, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Android com.kuangxiangciweimao.novel " + cat.version,
	}).HttpClient()
	for i := 0; i < 5; i++ {
		decodeText := DecodeText(response.String(), cat.decodeKey)
		if result := gjson.Parse(decodeText); result.Exists() {
			return result
		}
	}
	return gjson.Result{}
}

func (cat *Ciweimao) noDecodePost(url string, data map[string]string) gjson.Result {
	if data != nil {
		for k, v := range data {
			cat.params[k] = v
		}
	}
	response := buildHttps.Get(cat.host+url, cat.params, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Android com.kuangxiangciweimao.novel " + cat.version,
	}).HttpClient()
	for i := 0; i < 5; i++ {
		if result := gjson.Parse(response.String()); result.Exists() {
			return result
		}
	}
	return gjson.Result{}
}
