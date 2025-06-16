package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel-go/service"
)

// ResourceController 前台门户-资源(图片/视频/文档)模块控制器
type ResourceController struct {
	ResourceService *service.ResourceService
}

// NewResourceController 创建资源控制器实例
func NewResourceController(service *service.ResourceService) *ResourceController {
	return &ResourceController{
		ResourceService: service,
	}
}

// RegisterRoutes 注册资源相关路由
func (rc *ResourceController) RegisterRoutes(router *gin.RouterGroup) {
	resourceGroup := router.Group("/resource")
	{
		resourceGroup.GET("/img_verify_code", rc.GetImgVerifyCode)
		resourceGroup.POST("/image", rc.UploadImage)
	}
}

// GetImgVerifyCode GetImageVerifyCode 获取图片验证码
// @Summary 获取图片验证码接口
// @Description 返回图片验证码的 Base64 编码和对应的验证码内容
// @Tags Resource
// @Accept  json
// @Produce  json
// @Router /api/front/resource/img_verify_code [get]
func (rc *ResourceController) GetImgVerifyCode(c *gin.Context) {
	resp, err := rc.ResourceService.GetImgVerifyCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "生成验证码失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "00000", "msg": "成功", "data": resp})
}

// UploadImage 图片上传接口 Todo
func (rc *ResourceController) UploadImage(c *gin.Context) {
	token := c.GetHeader("Authorization")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "文件上传失败"})
		return
	}

	resp, err := rc.ResourceService.UploadImage(file, token, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "上传失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传成功", "data": resp})
}
