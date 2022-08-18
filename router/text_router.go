package router

import "github.com/gin-gonic/gin"

func LoadTestRouter(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}
