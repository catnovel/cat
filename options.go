package cat

type CiweimaoOption interface {
	apply(*Ciweimao)
}
type CiweimaoOptionFunc func(*Ciweimao)

func (optionFunc CiweimaoOptionFunc) apply(c *Ciweimao) {
	optionFunc(c)
}

func Debug() CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.debug = true
	})
}
func NoDecode() CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.decodeKey = ""
	})
}
func MaxRetry(retry int) CiweimaoOption {
	return CiweimaoOptionFunc(func(c *Ciweimao) {
		c.maxRetry = retry
	})
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
