package middlewares

import (
	"github.com/gin-gonic/gin"
	"im/helper"
	"net/http"
	"strings"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if strings.Contains(token, "bearer") {
			token = strings.ReplaceAll(token, "bearer ", "")
		}

		userClaims, err := helper.AnalyseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证不通过",
			})
			c.Abort()
			return
		}
		c.Set("user_claims", userClaims)
		c.Next()
	}
}
