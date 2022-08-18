package middleware

import (
	"bookmanager/model"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用户携带Token 的方式：1、放在请求头  2、放在请求体  3、放在url
		// token验证成功，返回c.Next()继续，否则返回c.Abort()直接返回
		token := c.Request.Header.Get("token")
		u := model.User{}
		if token == "" {
			c.String(401, "Token无效或者未携带Token")
			c.Abort()
			return
		}
		c.Set("UserId", u.Id)
		c.Next()
	}
}
