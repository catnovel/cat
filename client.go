package cat

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/catnovelapi/BuilderHttpClient"
	"github.com/tidwall/gjson"
	"log"
)

type Ciweimao struct {
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
	} else {
		LoginToken(loginToken).apply(cat)
		Account(account).apply(cat)
	}
}

func (cat *Ciweimao) post(url string, data map[string]any, options ...HttpOption) gjson.Result {
	httpBuilder := &HttpClient{maxRetry: 3, debug: false, decodeKey: cat.decodeKey}
	for _, option := range options {
		option.apply(httpBuilder)
	}
	for i := 0; i < httpBuilder.maxRetry; i++ {
		response := BuilderHttpClient.Post(cat.host+url, BuilderHttpClient.Body(cat.setParams(data)), BuilderHttpClient.Header(cat.headers))
		if httpBuilder.debug {
			response = response.Debug()
		}
		resultText, err := cat.DecodeEncryptText(response.Text(), httpBuilder.decodeKey)
		if err == nil {
			return gjson.Parse(resultText)
		} else {
			log.Println("解密失败 ", err)
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

// LoadKey 读取解密密钥
func LoadKey(EncryptKey string) []byte {
	Key := SHA256([]byte(EncryptKey))
	return Key[:32]
}

func aesDecrypt(EncryptKey string, ciphertext []byte) ([]byte, error) {
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

func (cat *Ciweimao) DecodeEncryptText(str string, decodeKey string) (string, error) {
	if decodeKey == "" {
		return str, nil
	}
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	raw, err := aesDecrypt(decodeKey, decoded)
	if err != nil {
		return "", err
	}
	if len(raw) == 0 {
		return "", errors.New("解密内容为空,请检查解密内容内容是否正确")
	}
	return string(raw), nil
}
