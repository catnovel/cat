package cat

type HttpClient struct {
	debug     bool
	maxRetry  int
	decodeKey string
}

type HttpOption interface {
	apply(*HttpClient)
}
type HttpOptionFunc func(*HttpClient)

func (optionFunc HttpOptionFunc) apply(c *HttpClient) {
	optionFunc(c)
}
func Debug() HttpOption {
	return HttpOptionFunc(func(c *HttpClient) {
		c.debug = true
	})
}
func NoDecode() HttpOption {
	return HttpOptionFunc(func(c *HttpClient) {
		c.decodeKey = ""
	})
}
func MaxRetry(retry int) HttpOption {
	return HttpOptionFunc(func(c *HttpClient) {
		c.maxRetry = retry
	})
}

type CiweimaoOption interface {
	apply(*Ciweimao)
}
type CiweimaoOptionFunc func(*Ciweimao)

func (optionFunc CiweimaoOptionFunc) apply(c *Ciweimao) {
	optionFunc(c)
}

func ApiBase(host string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.host = host
	})
}
func Version(version string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.version = version
	})
}
func DecodeKey(decodeKey string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.decodeKey = decodeKey
	})
}
func DeviceToken(deviceToken string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.deviceToken = deviceToken
	})
}
func AppVersion(appVersion string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.version = appVersion
	})
}
func LoginToken(loginToken string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.loginToken = loginToken
	})
}
func Account(account string) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.account = account
	})
}
