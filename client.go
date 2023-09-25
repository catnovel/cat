package cat

import (
	"github.com/catnovelapi/BuilderHttpClient"
	"github.com/tidwall/gjson"
	"log"
)

type Ciweimao struct {
	debug       bool
	maxRetry    int
	host        string
	version     string
	loginToken  string
	account     string
	deviceToken string
	decodeKey   string
	headers     map[string]any
}

func NewCiweimaoClient(options ...CiweimaoOption) *Ciweimao {
	client := &Ciweimao{
		maxRetry:    3,
		host:        "https://app.hbooker.com",
		version:     "2.9.290",
		decodeKey:   "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn",
		deviceToken: "ciweimao_",
	}
	for _, option := range options {
		option.apply(client)
	}

	client.headers = map[string]any{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Android com.kuangxiangciweimao.novel " + client.version,
	}

	return client
}

func (cat *Ciweimao) setParams(data map[string]any) map[string]any {
	params := map[string]any{
		"device_token": cat.deviceToken,
		"app_version":  cat.version,
		"login_token":  cat.loginToken,
		"account":      cat.account,
	}
	if data != nil {
		for k, v := range data {
			params[k] = v
		}
	}
	return params

}
func (cat *Ciweimao) NewAuthentication(loginToken, account string) {
	if len(loginToken) != 32 {
		log.Printf("loginToken长度不正确,必须为32位,当前变量:%s", loginToken)
		return
	}
	cat.loginToken = loginToken
	cat.account = account

}
func (cat *Ciweimao) post(url string, data map[string]any, options ...CiweimaoOption) gjson.Result {
	for _, option := range options {
		option.apply(cat)
	}
	for i := 0; i < cat.maxRetry; i++ {

		response := BuilderHttpClient.Post(cat.host+url, BuilderHttpClient.Body(cat.setParams(data)), BuilderHttpClient.Header(cat.headers))

		if cat.debug {
			response = response.Debug()
		}
		var resultText string
		if cat.decodeKey == "" {
			resultText = response.Text()
		} else {
			resultText = DecodeText(response.Text(), cat.decodeKey)
		}
		if result := gjson.Parse(resultText); result.Exists() {
			return result
		}
	}
	return gjson.Result{}
}
