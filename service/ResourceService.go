package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"mime/multipart"
	"novel-go/config"
	"novel-go/constants"
	"novel-go/model/resp"
	"time"
)

// ResourceService 资源服务接口
type ResourceService interface {
	// GetImgVerifyCode 获取图片验证码
	// @Summary 获取图片验证码
	// @Description 生成数字图片验证码，返回 base64 和 UUID
	// @Tags 资源模块
	// @Produce json
	// @Success 200 {object} resp.ImgVerifyCodeRespDto
	// @Failure 500 {string} string "生成验证码失败"
	// @Router /resource/img-verify-code [get]
	GetImgVerifyCode() (*resp.ImgVerifyCodeRespDto, error)

	// UploadImage 上传图片
	// @Summary 上传图片
	// @Description 接收图片文件上传，返回图片路径
	// @Tags 资源模块
	// @Accept multipart/form-data
	// @Produce json
	// @Param file formData file true "图片文件"
	// @Param token header string true "用户Token"
	// @Success 200 {string} string "图片保存路径"
	// @Failure 400 {string} string "上传失败"
	// @Router /resource/upload-image [post]
	UploadImage(file *multipart.FileHeader, token string, c *gin.Context) (string, error)
}

// ResourceServiceImpl 资源服务实现
type ResourceServiceImpl struct {
}

func NewResourceServiceImpl() ResourceService {
	return &ResourceServiceImpl{}
}

func (s *ResourceServiceImpl) GetImgVerifyCode() (*resp.ImgVerifyCodeRespDto, error) {
	captcha := base64Captcha.NewCaptcha(
		base64Captcha.NewDriverDigit(45, 120, 5, 0.7, 80),
		base64Captcha.DefaultMemStore,
	)
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return nil, err
	}
	err = config.RedisClient.Set(config.Ctx, constants.RedisKeyCaptcha+id, answer, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return &resp.ImgVerifyCodeRespDto{
		Img:  b64s,
		UUID: id,
	}, nil
}

func (s *ResourceServiceImpl) UploadImage(file *multipart.FileHeader, token string, c *gin.Context) (string, error) {
	dst := "uploads/" + file.Filename
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		return "", err
	}
	return dst, nil
}
