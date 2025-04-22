package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hhdms/msjx/internal/models"
	"github.com/hhdms/msjx/internal/utils"
)

// JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取令牌
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 解析令牌
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
				Code: 0,
				Msg:  "无效的令牌",
				Data: nil,
			})
			return
		}

		// 将用户信息保存到上下文
		c.Set("userId", claims.ID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
