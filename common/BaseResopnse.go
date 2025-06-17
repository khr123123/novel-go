package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // 没有数据时可省略
}

// SuccessResponse 统一成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:    "00000",
		Message: "success",
		Data:    data,
	})
}

// ErrorResponse 统一失败响应
func ErrorResponse(c *gin.Context, code string, message string) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:    code,
		Message: message,
	})
}
