package test

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	captcha := base64Captcha.NewCaptcha(base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80), base64Captcha.DefaultMemStore)
	// 生成验证码
	id, b64s, answer, err := captcha.Generate()
	fmt.Println(id)
	fmt.Println(b64s)

	fmt.Println(answer)

	fmt.Println(err)

}
