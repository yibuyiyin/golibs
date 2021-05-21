package captcha

import "github.com/mojocn/base64Captcha"

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

//var driver base64Captcha.Driver
var store = rdsMemStore
var driver = base64Captcha.DefaultDriverDigit

var err error

func Generate() (id, b64s string) {
	c := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
	id, b64s, err = c.Generate()
	if err != nil {
		panic(err)
	}
	return
}

func Verify(id, answer string) bool {
	c := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
	return c.Verify(id, answer, true)
}
