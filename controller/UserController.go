package controller

import (
	"net/http"
	"novel-go/model/req"
	"novel-go/service"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	UserService *service.UserService
}

// NewUserController 创建UserController实例
func NewUserController(userService *service.UserService /**bookService *service.BookService**/) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// RegisterRoutes 注册用户相关路由
func (uc *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/user")
	{
		userGroup.POST("/register", uc.Register)
		//userGroup.POST("/login", uc.Login)
	}
}

// Register 用户注册接口
// @Summary 用户注册
// @Description 用户通过用户名、密码和验证码注册账号
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body req.UserRegisterReqDto true "注册请求参数"
// @Success 200 {object} gin.H{code=int,msg=string,data=resp.UserRegisterRespDto} "注册成功"
// @Failure 400 {object} gin.H{code=int,msg=string} "请求参数错误"
// @Failure 500 {object} gin.H{code=int,msg=string} "服务器内部错误"
// @Router /api/front/user/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var reqDto req.UserRegisterReqDto
	if err := c.ShouldBindJSON(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	resp, err := uc.UserService.Register(&reqDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功", "data": resp})
}

// Login 用户登录接口
//func (uc *UserController) Login(c *gin.Context) {
//	var req req.UserLoginReqDto
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
//		return
//	}
//	resp, err := uc.UserService.Login(&req)
//	if err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": resp})
//}
