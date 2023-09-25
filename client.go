package cat

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
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
	App         *CiweimaoApp
}

type CiweimaoApp struct {
	api *Ciweimao
}

func NewCiweimaoClient(options ...CiweimaoOption) *Ciweimao {
	client := &Ciweimao{
		maxRetry:    3,
		host:        "https://app.hbooker.com",
		version:     "2.9.290",
		decodeKey:   "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn",
		deviceToken: "ciweimao_",
		//App:         &CiweimaoApp{},
	}
	for _, option := range options {
		option.apply(client)
	}

	client.headers = map[string]any{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Android com.kuangxiangciweimao.novel " + client.version,
	}

	client.App = &CiweimaoApp{api: client}

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

var IV = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SHA256 sha256 编码
func SHA256(data []byte) []byte {
	ret := sha256.Sum256(data)
	return ret[:]
}

// Base64Decode Base64 解码
func Base64Decode(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

// LoadKey 读取解密密钥
func LoadKey(EncryptKey string) []byte {
	Key := SHA256([]byte(EncryptKey))
	return Key[:32]
}

// AESDecrypt AES 解密
func AESDecrypt(EncryptKey string, ciphertext []byte) ([]byte, error) {
	key := LoadKey(EncryptKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, IV)
	plainText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plainText, ciphertext)
	plainText = PKCS7UnPadding(plainText)
	return plainText, nil
}

// PKCS7UnPadding 对齐
func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unpadding := int(plainText[length-1])
	return plainText[:(length - unpadding)]
}

// DecodeText 入口函数
func DecodeText(str string, encryptKey string) string {
	var err error
	var decoded, raw []byte
	decoded, err = Base64Decode(str)
	if err != nil {
		return str
	}
	raw, err = AESDecrypt(encryptKey, decoded)
	if err != nil {
		panic(err)
	}
	return string(raw)
}
