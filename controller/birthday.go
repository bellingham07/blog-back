package controller

import (
	"blog-back/database"
	"blog-back/model"

	"github.com/gin-gonic/gin"
)

func RegisterBirthday(c *gin.Context) {
	db := database.GetDB()

	birth := c.PostForm("birthday")
	email := c.PostForm("email")
	// 校验邮箱和生日是否合法

	// 保存到数据库
	db.Save(&model.Birthday{
		Name:     email,
		Birthday: birth,
		Email:    email,
	})

}
