package router

import (
	"blog-back/controller"

	"github.com/gin-gonic/gin"
)

type BirthdayRoute struct {
}

func (*BirthdayRoute) InitRouter(g *gin.RouterGroup) {
	u := g.Group("/birthday")
	// u.Use(middleware.RefreshTokenMiddleware())
	// 不需要拦截器
	{
		u.POST("/register", controller.RegisterBirthday)

	}
	// 需要拦截器
	// c := g.Group("/user")
	// c.Use(middleware.RefreshTokenMiddleware(), middleware.AuthMiddleware())
	// {
	// 	c.POST("/id", controller.QueryById)
	// }
}
