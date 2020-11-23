package main

import (
	"github.com/zhaoyunxing92/letter/conf"
	"github.com/zhaoyunxing92/letter/database"
	"github.com/zhaoyunxing92/letter/middleware"
	_ "github.com/zhaoyunxing92/letter/global"
	"github.com/zhaoyunxing92/letter/router"
	"log"
)

func main() {
	//初始化配置
	conf.Setup("application.yml")
	//初始化mongo
	database.Setup()
	//启动路由
	engine := router.InitRouter(middleware.TranslationMiddleware())
	if err := engine.Run(); err != nil {
		log.Panicf("letter run error:%v", err)
	}
}
