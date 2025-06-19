package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"novel-go/common"
	"novel-go/model/req"
	"novel-go/service"
)

type SearchController struct {
	SearchService service.SearchService
}

func NewSearchController(s service.SearchService) *SearchController {
	return &SearchController{SearchService: s}
}
func (c *SearchController) RegisterRoutes(r *gin.RouterGroup) {
	searchGroup := r.Group("/search")
	{
		searchGroup.GET("/books", c.SearchBooks)
	}
}

func (c *SearchController) SearchBooks(ctx *gin.Context) {
	var req req.BookSearchReqDto
	if err := ctx.ShouldBindQuery(&req); err != nil {
		common.ErrorResponse(ctx, "4000", err.Error())
		return
	}
	// 打印绑定后的请求结构体，方便调试
	fmt.Printf("请求参数: %+v\n", req)

	result, err := c.SearchService.SearchBooks(req)
	if err != nil {
		common.ErrorResponse(ctx, "5000", err.Error())
		return
	}
	common.SuccessResponse(ctx, result)
}
