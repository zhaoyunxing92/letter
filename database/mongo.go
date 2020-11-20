package database

import (
	"context"
	"github.com/zhaoyunxing92/letter/conf"
	"github.com/zhaoyunxing92/letter/public"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// Setup:初始化数据库
func Setup() {
	cfg := conf.MongodbConfig

	clientOptions := options.Client().ApplyURI(cfg.Url).SetMaxPoolSize(cfg.MaxPoolSize).SetMinPoolSize(cfg.MinPoolSize).SetConnectTimeout(time.Duration(cfg.Timeout) * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("connect mogo fail: %s", err.Error())
	}
	//检查是否ping通
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatalf("mogo ping fail: %s", err.Error())
	}
	//设置数据库实例
	public.Mongo = client.Database(cfg.Name)
}