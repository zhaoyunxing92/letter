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
	//applet:=domain.AppletInfo{}
	//collection := applet.GetCollection()

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//_, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	//if err!=nil{
	//	log.Fatalf("collection fail:%s",err.Error())
	//}
	//启动路由
	engine := router.InitRouter(middleware.TranslationMiddleware())
	if err := engine.Run(); err != nil {
		log.Panicf("letter run error:%v", err)
	}
}
