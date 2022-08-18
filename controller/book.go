package controller

import (
	"bookmanager/dao/mysql"
	"bookmanager/model"

	"github.com/gin-gonic/gin"
)

// 创建书籍添加数据
func CreateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	mysql.DB.Create(p)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}

// 查看书籍列表
func GetBookListHandler(c *gin.Context) {
	books := []model.Book{}
	mysql.DB.Find(&books)
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": books,
	})
}

// 查看单个书籍
func GetBookHandler(c *gin.Context) {
	id := c.Param("id")
	book := model.Book{}
	mysql.DB.Find(&book, id)
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": book,
	})
}

// 修改书籍
func UpdateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	oldBook := &model.Book{Id: p.Id}
	var newBook model.Book
	if p.Name != "" {
		newBook.Name = p.Name
	}
	if p.Desc != "" {
		newBook.Desc = p.Desc
	}
	mysql.DB.Model(oldBook).Updates(newBook)
	c.JSON(200, gin.H{
		"msg":     "success",
		"newBook": newBook,
	})
}

// 删除书籍
func DeleteBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	mysql.DB.Select("Users").Delete(model.Book{Id: p.Id})
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
