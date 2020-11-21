package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/letter/ctrl"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	router.Use(middlewares...)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	v1 := router.Group("/v1")
	//应用创建模块
	applet := v1.Group("/applet")
	{ //创建应用
		applet.POST("/", ctrl.CreateApplet)
	}
	//模板模块
	template := v1.Group("/template")
	{ //创建模板
		template.POST("/")
	}

	return router
}
