package controller

import (
	"bookmanager/dao/mysql"
	"bookmanager/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterHandler(c *gin.Context) {
	p := new(model.User)
	// 参数校验与绑定
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	//，创建数据
	mysql.DB.Create(p)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func LoginHandler(c *gin.Context) {
	p := new(model.User)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// 查询数据
	u := model.User{Username: p.Username, Password: p.Password}
	if rows := mysql.DB.Where(&u).First(&u).Row(); rows == nil {
		c.JSON(400, gin.H{
			"message": "用户名或密码错误",
		})
		return
	}

	// 随机生成一个字符串作为Token
	token := uuid.New().String()
	// 将Token存入数据库
	mysql.DB.Model(&u).Update("token", token)
	c.JSON(200, gin.H{
		"message": "success",
		"token":   token,
	})
}
