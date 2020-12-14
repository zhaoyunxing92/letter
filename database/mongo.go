package database

import (
	"context"
	"github.com/zhaoyunxing92/letter/conf"
	"github.com/zhaoyunxing92/letter/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// Setup:初始化数据库
func Setup() {
	cfg := conf.MongodbConfig
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	clientOptions := options.Client().ApplyURI(cfg.Url).SetMaxPoolSize(cfg.MaxPoolSize).SetMinPoolSize(cfg.MinPoolSize).SetConnectTimeout(time.Duration(cfg.Timeout) * time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("connect mogo fail: %s", err.Error())
	}
	//检查是否ping通
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("mogo ping fail: %s", err.Error())
	}
	//设置数据库实例
	global.Mongo = client.Database(cfg.Name)
}
