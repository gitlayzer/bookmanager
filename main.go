package main

import (
	"bookmanager/dao/mysql"
	"bookmanager/router"
)

func main() {
	// 初始化MySQL
	mysql.InitMysql()
	// 将实例化router服务的方法拆分到项目的router沐浴露下
	r := router.InitRouter()
	r.Run(":80")
}
