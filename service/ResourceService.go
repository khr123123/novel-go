package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"mime/multipart"
	"novel-go/config"
	"novel-go/constants"
	"novel-go/model/req"
	"time"
)

type ResourceService struct{}

func NewResourceService() *ResourceService {
	return &ResourceService{}
}

func (s *ResourceService) GetImgVerifyCode() (*req.ImgVerifyCodeRespDto, error) {
	// 这里使用 captcha 包生成验证码
	captcha := base64Captcha.NewCaptcha(
		base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80),
		base64Captcha.DefaultMemStore,
	)
	id, b64s, answer, _ := captcha.Generate()
	// 把答案存入 Redis，设置过期时间 5 分钟
	err := config.RedisClient.Set(config.Ctx, constants.RedisKeyCaptcha+id, answer, 5*time.Minute).Err()
	if err != nil {
		return nil, nil
	}
	return &req.ImgVerifyCodeRespDto{
		Img:  b64s,
		UUID: id,
	}, nil
}

func (s *ResourceService) UploadImage(file *multipart.FileHeader, token string, c *gin.Context) (string, error) {
	// 保存图片或上传到 OSS、MinIO 等对象存储
	dst := "uploads/" + file.Filename
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		return "", err
	}
	return dst, nil
}
