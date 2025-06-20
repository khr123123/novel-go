package controller

import (
	"novel-go/common"
	"novel-go/config"
	"novel-go/model/pojo"
	"novel-go/model/req"
	"novel-go/model/resp"
	"novel-go/service"
	"novel-go/utils"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	UserService service.UserService
	BookService service.BookService
}

// NewUserController 创建UserController实例
func NewUserController(userService service.UserService, bookService service.BookService) *UserController {
	return &UserController{
		UserService: userService,
		BookService: bookService,
	}
}

// RegisterRoutes 注册用户相关路由
func (uc *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/user")
	{
		userGroup.POST("/register", uc.Register)
		userGroup.POST("/login", uc.Login)
		userGroup.POST("/comment", uc.Comment)
		userGroup.GET("", uc.GetUserInfo)
	}
}

// Register 用户注册接口
// @Summary 用户注册
// @Description 用户通过用户名、密码和验证码注册账号
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body req.UserRegisterReqDto true "注册请求参数"
// @Success 200 {object} BaseResponse{code=int,message=string,data=resp.UserRegisterRespDto} "注册成功"
// @Failure 400 {object} BaseResponse{code=int,message=string} "请求参数错误"
// @Failure 500 {object} BaseResponse{code=int,message=string} "服务器内部错误"
// @Router /api/front/user/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var reqDto req.UserRegisterReqDto
	if err := c.ShouldBindJSON(&reqDto); err != nil {
		common.ErrorResponse(c, "4000", err.Error())
		return
	}
	resp, err := uc.UserService.Register(&reqDto)
	if err != nil {
		common.ErrorResponse(c, "5000", err.Error())
		return
	}
	common.SuccessResponse(c, resp)
}

// Login 用户登录接口
func (uc *UserController) Login(c *gin.Context) {
	var loginReq req.UserLoginReqDto
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		common.ErrorResponse(c, "4000", err.Error())
		return
	}
	resp, err := uc.UserService.Login(&loginReq)
	if err != nil {
		common.ErrorResponse(c, "4001", err.Error())
		return
	}
	common.SuccessResponse(c, resp)
}

func (uc *UserController) Comment(c *gin.Context) {
	id, _ := utils.GetUserID(c.Request.Context())
	var commentReqDto req.UserCommentReqDto
	if err := c.ShouldBindJSON(&commentReqDto); err != nil {
		common.ErrorResponse(c, "4000", err.Error())
		return
	}
	commentReqDto.UserId = id
	err := uc.BookService.SaveComment(c, commentReqDto)
	if err != nil {
		common.ErrorResponse(c, "5000", err.Error())
		return
	}
	common.SuccessResponse(c, nil)
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	id, _ := utils.GetUserID(c.Request.Context())
	var userinfo pojo.UserInfo
	err := config.DB.WithContext(c).
		Where("id = ?", id).
		First(&userinfo).Error
	if err != nil {
		common.ErrorResponse(c, "4000", err.Error())
		return
	}
	common.SuccessResponse(c, resp.UserInfoRespDto{
		NickName:  userinfo.NickName,
		UserPhoto: userinfo.UserPhoto,
		UserSex:   userinfo.UserSex,
	})
}
