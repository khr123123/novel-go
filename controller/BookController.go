package controller

import (
	"novel-go/common"
	"novel-go/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BookController 小说模块控制器
type BookController struct {
	BookService service.BookService
}

// NewBookController 创建 BookController 实例
func NewBookController(bookService service.BookService) *BookController {
	return &BookController{
		BookService: bookService,
	}
}

// RegisterRoutes 注册路由
func (c *BookController) RegisterRoutes(rg *gin.RouterGroup) {
	books := rg.Group("/book")
	{
		books.GET("/category/list", c.ListCategory)
		books.GET("/:id", c.GetBookByID)
		books.POST("/visit", c.AddVisitCount)
		books.GET("/last_chapter/about", c.GetLastChapterAbout)
		books.GET("/rec_list", c.ListRecBooks)
		books.GET("/chapter/list", c.ListChapters)
		books.GET("/content/:chapterId", c.GetBookContentAbout)
		books.GET("/pre_chapter_id/:chapterId", c.GetPreChapterID)
		books.GET("/next_chapter_id/:chapterId", c.GetNextChapterID)
		books.GET("/visit_rank", c.ListVisitRankBooks)
		books.GET("/newest_rank", c.ListNewestRankBooks)
		books.GET("/update_rank", c.ListUpdateRankBooks)
		books.GET("/comment/newest_list", c.ListNewestComments)
	}
}

// ListCategory 小说分类列表查询接口
func (c *BookController) ListCategory(ctx *gin.Context) {
	workDirectionStr := ctx.Query("workDirection")
	workDirection, err := strconv.Atoi(workDirectionStr)
	if err != nil {
		common.ErrorResponse(ctx, "40000", "workDirection 参数错误")
		return
	}

	data, err := c.BookService.ListCategory(ctx, workDirection)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// GetBookByID 小说信息查询接口
func (c *BookController) GetBookByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	data, err := c.BookService.GetBookById(ctx, idStr)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

type AddVisitCountReq struct {
	BookId string `json:"bookId"`
}

// AddVisitCount 增加小说点击量接口
func (c *BookController) AddVisitCount(ctx *gin.Context) {
	var req AddVisitCountReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ErrorResponse(ctx, "400", "参数绑定失败")
		return
	}
	err := c.BookService.AddVisitCount(ctx, req.BookId)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, nil)
}

// GetLastChapterAbout 小说最新章节相关信息查询接口
func (c *BookController) GetLastChapterAbout(ctx *gin.Context) {
	bookIdStr := ctx.Query("bookId")
	data, err := c.BookService.GetLastChapterAbout(ctx, bookIdStr)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// ListRecBooks 小说推荐列表查询接口
func (c *BookController) ListRecBooks(ctx *gin.Context) {
	bookIdStr := ctx.Query("bookId")
	data, err := c.BookService.ListRecBooks(ctx, bookIdStr)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// ListChapters 小说章节列表查询接口
func (c *BookController) ListChapters(ctx *gin.Context) {
	bookIdStr := ctx.Query("bookId")
	data, err := c.BookService.ListChapters(ctx, bookIdStr)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// GetBookContentAbout 小说内容相关信息查询接口
func (c *BookController) GetBookContentAbout(ctx *gin.Context) {
	chapterIdStr := ctx.Param("chapterId")
	data, err := c.BookService.GetBookContentAbout(ctx, chapterIdStr)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// GetPreChapterID 获取上一章节ID接口
func (c *BookController) GetPreChapterID(ctx *gin.Context) {
	chapterIdStr := ctx.Param("chapterId")
	data, err := c.BookService.GetPreChapterId(ctx, chapterIdStr)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// GetNextChapterID 获取下一章节ID接口
func (c *BookController) GetNextChapterID(ctx *gin.Context) {
	chapterIdStr := ctx.Param("chapterId")
	data, err := c.BookService.GetNextChapterId(ctx, chapterIdStr)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// ListVisitRankBooks 小说点击榜查询接口
func (c *BookController) ListVisitRankBooks(ctx *gin.Context) {
	data, err := c.BookService.ListVisitRankBooks(ctx)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// ListNewestRankBooks 小说新书榜查询接口
func (c *BookController) ListNewestRankBooks(ctx *gin.Context) {
	data, err := c.BookService.ListNewestRankBooks(ctx)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// ListUpdateRankBooks 小说更新榜查询接口
func (c *BookController) ListUpdateRankBooks(ctx *gin.Context) {
	data, err := c.BookService.ListUpdateRankBooks(ctx)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}

// ListNewestComments 小说最新评论查询接口
func (c *BookController) ListNewestComments(ctx *gin.Context) {
	bookIdStr := ctx.Query("bookId")
	data, err := c.BookService.ListNewestComments(ctx, bookIdStr)
	if err != nil {
		common.ErrorResponse(ctx, "50000", err.Error())
		return
	}
	common.SuccessResponse(ctx, data)
}
