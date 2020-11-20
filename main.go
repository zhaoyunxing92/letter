package main

import (
	"context"
	"github.com/zhaoyunxing92/letter/conf"
	"github.com/zhaoyunxing92/letter/database"
	"github.com/zhaoyunxing92/letter/domain"
	_ "github.com/zhaoyunxing92/letter/public"
	"github.com/zhaoyunxing92/letter/router"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func main() {
	//初始化配置
	conf.Setup("application.yml")
	//初始化mongo
	database.Setup()
	applet:=domain.Applet{}
	collection := applet.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err!=nil{
		log.Fatalf("collection fail:%s",err.Error())
	}
	//启动路由
	engine := router.InitRouter()
	if err := engine.Run(); err != nil {
		panic("gin error")
	}
}
