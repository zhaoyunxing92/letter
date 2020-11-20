package main

import (
	"github.com/zhaoyunxing92/letter/conf"
	"github.com/zhaoyunxing92/letter/database"
	_ "github.com/zhaoyunxing92/letter/public"
	"github.com/zhaoyunxing92/letter/router"
)

func main() {
	//初始化配置
	conf.Setup("application.yml")
	//初始化mongo
	database.Setup()
	//启动路由
	engine := router.InitRouter()
	if err := engine.Run(); err != nil {
		panic("gin error")
	}
}
