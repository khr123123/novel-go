package controller

import (
	"github.com/gin-gonic/gin"
	"novel-go/common"
	"novel-go/service"
)

// ResourceController 前台门户-资源(图片/视频/文档)模块控制器
type ResourceController struct {
	ResourceService service.ResourceService
}

// NewResourceController 创建资源控制器实例
func NewResourceController(service service.ResourceService) *ResourceController {
	return &ResourceController{
		ResourceService: service,
	}
}

// RegisterRoutes 注册资源相关路由
func (rc *ResourceController) RegisterRoutes(router *gin.RouterGroup) {
	resourceGroup := router.Group("/resource")
	{
		resourceGroup.GET("/img_verify_code", rc.GetImgVerifyCode)
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
		common.ErrorResponse(c, "5000", "生成验证码失败,请稍后重试..")
		return
	}
	common.SuccessResponse(c, resp)
}
