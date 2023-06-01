package router

import (
	"github.com/gin-gonic/gin"
	v1 "test_go_mod/api/v1"
	"test_go_mod/controller"
	"test_go_mod/tool/validator"
)

func Http(route *gin.Engine) {
	var apiRouter = route.Group("/api/v1")
	{
		apiRouter.POST("/detect", v1.Detect)
		apiRouter.GET("/resp", v1.Resp)
		apiRouter.POST("/login", v1.Login)
		//apiRouter.POST("/register", validator.GetHandlerFunc(v1.Register))
	}

	var authRouter = route.Group("/api/v1/auth")
	{
		authRouter.POST("/login", v1.Login)
		authRouter.POST("/register", validator.GetHandlerFunc(v1.Register))
	}

	route.GET("list", controller.GetUserList)
	route.GET("user", controller.GetUser)
	//route.GET("create_user", controller.CreateUser)
	route.GET("test", controller.Test)

	//route.GET("/test", func(c *gin.Context) {
	//	developer := c.Query("name")
	//	c.JSON(200, gin.H{
	//		"message":   "人人租机",
	//		"url":       "https://www.rrzu.com",
	//		"developer": developer,
	//	})
	//})
}
