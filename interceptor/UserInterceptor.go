package interceptor

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"novel-go/common"
	"novel-go/config"
	"novel-go/constants"
	"novel-go/model/pojo"
	"novel-go/utils"
)

// 白名单路径前缀，访问时不需要鉴权
var whitelistPrefixes = []string{
	"/api/front/user/register",
	"/api/front/user/login",
	"/api/admin/login",

	"/api/front/book",
	"/api/front/home",
	"/api/front/news",
	"/api/front/resource",
	"/api/front/search",
}

func UserInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// 1. 路径白名单，直接放行
		for _, prefix := range whitelistPrefixes {
			if strings.HasPrefix(path, prefix) {
				c.Next()
				return
			}
		}

		// 2. 获取 Header 中 token
		token := c.GetHeader(constants.HttpAuthHeaderName)
		if token == "" {
			common.ErrorResponse(c, "A0230", "用户登录已过期或未登录")
			c.Abort()
			return
		}

		// 3. 解析 token 得到 userId
		userId, err := utils.ParseToken(token, constants.JwtSecret)
		if err != nil || userId <= 0 {
			common.ErrorResponse(c, "A0230", "用户登录已过期或未登录")
			c.Abort()
			return
		}

		// 4. 数据库验证 userId 是否有效
		var count int64
		err = config.DB.WithContext(c).Model(&pojo.UserInfo{}).Where("id = ?", userId).Count(&count).Error
		if err != nil {
			common.ErrorResponse(c, "A0230", "数据库查询失败")
			c.Abort()
			return
		}
		if count != 1 {
			common.ErrorResponse(c, "A0230", "用户不存在或已被禁用")
			c.Abort()
			return
		}
		fmt.Print("userId", userId)
		// 5. 设置 userId 到请求上下文中，方便后续业务使用
		ctx := utils.SetUserID(c.Request.Context(), userId)
		c.Request = c.Request.WithContext(ctx)

		// 继续请求处理
		c.Next()
	}
}
