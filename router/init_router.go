package router

import "github.com/gin-gonic/gin"

// 此方法作用是初始化其他文件的路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	LoadTestRouter(r)
	LoadApiRouter(r)
	return r
}
