package controller

import (
	"github.com/gin-gonic/gin"
	"novel-go/common"
	"novel-go/service"
)

// HomeController 前台门户-首页模块
type HomeController struct {
	HomeService service.HomeService
}

// NewHomeController 构造函数
func NewHomeController(homeService service.HomeService) *HomeController {
	return &HomeController{
		HomeService: homeService,
	}
}

// RegisterRoutes 路由注册
func (h *HomeController) RegisterRoutes(router *gin.RouterGroup) {
	homeGroup := router.Group("/home")
	{
		homeGroup.GET("/books", h.ListHomeBooks)
		homeGroup.GET("/friend_Link/list", h.ListHomeFriendLinks)
	}
}

// ListHomeBooks 首页小说推荐查询接口
func (h *HomeController) ListHomeBooks(c *gin.Context) {
	books, err := h.HomeService.ListHomeBooks(c)
	if err != nil {
		common.ErrorResponse(c, "5000", "查询小说推荐失败")
		return
	}
	common.SuccessResponse(c, books)
}

// ListHomeFriendLinks 首页友情链接列表查询接口
func (h *HomeController) ListHomeFriendLinks(c *gin.Context) {
	links, err := h.HomeService.ListHomeFriendLinks(c)
	if err != nil {
		common.ErrorResponse(c, "5000", "查询友情链接失败")
		return
	}
	common.SuccessResponse(c, links)
}
