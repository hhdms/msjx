package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hhdms/msjx/internal/models"
	"github.com/hhdms/msjx/internal/services"
)

// AuthController 认证控制器
type AuthController struct {
	empService services.EmpService
}

// NewAuthController 创建认证控制器实例
func NewAuthController() *AuthController {
	return &AuthController{
		empService: services.NewEmpService(),
	}
}

// Login 处理登录请求
func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Code: 0,
			Msg:  "无效的请求参数",
			Data: nil,
		})
		return
	}

	// 调用登录服务
	resp, err := c.empService.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, models.Response{
		Code: 1,
		Msg:  "success",
		Data: resp,
	})
}
