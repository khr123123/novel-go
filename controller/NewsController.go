package controller

import (
	"novel-go/common"
	"novel-go/config"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"novel-go/service"
)

// NewsController 前台门户 - 新闻模块
type NewsController struct {
	NewsService service.NewsService
}

// NewNewsController 创建 NewsController 实例
func NewNewsController(newsService service.NewsService) *NewsController {
	return &NewsController{NewsService: newsService}
}

// RegisterRoutes 注册路由
func (c *NewsController) RegisterRoutes(rg *gin.RouterGroup) {
	news := rg.Group("/news")
	{
		news.GET("/latest_list", config.RedisCacheHandler(c.ListLatestNews, time.Minute)) // 最新新闻列表
		news.GET("/:id", config.RedisCacheHandler(c.GetNews, time.Minute))                // 查询新闻详情
	}
}

// ListLatestNews 最新新闻列表查询接口
func (c *NewsController) ListLatestNews(ctx *gin.Context) {
	newsList, err := c.NewsService.ListLatestNews(ctx)
	if err != nil {
		common.ErrorResponse(ctx, "50001", "获取最新新闻失败")
		return
	}
	common.SuccessResponse(ctx, newsList)
}

// GetNews 新闻信息查询接口
func (c *NewsController) GetNews(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ErrorResponse(ctx, "40001", "无效的新闻ID")
		return
	}
	news, err := c.NewsService.GetNews(ctx, id)
	if err != nil {
		common.ErrorResponse(ctx, "50002", "获取新闻信息失败")
		return
	}
	common.SuccessResponse(ctx, news)
}
