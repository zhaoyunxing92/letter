package domain

import (
	"context"
	"github.com/zhaoyunxing92/letter/global"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Applet struct {
	CorpId         string `json:"corpId"`            //企业id
	AgentId        string `json:"agentId"`           //应用id
	AppKey         string `json:"appKey" bson:"_id"` //应用key
	AppSecret      string `json:"appSecret"`         //应用秘钥
	Name           string `json:"name"`              //应用名称
	Icon           string `json:"icon"`              //应用图标
	Desc           string `json:"desc"`              //应用描述
	Self           bool   `json:"self"`              //是否是自建应用false：不是,true：自建
	OmpLink        string `json:"ompLink"`           //应用的OA后台管理主页
	HomepageLink   string `json:"homepageLink"`      //应用的移动端主页。
	PcHomepageLink string `json:"pcHomepageLink"`    //应用的PC端主页。
	Status         uint8  `json:"status"`            //应用状态
	Version        uint   `json:"version"`           //数据版本
	CreateIn       int64  `json:"createIn"`          //创建时间
	UpdateIn       int64  `json:"updateIn"`          //创建时间
}

//获取集合
func (applet *Applet) collection() *mongo.Collection {
	return global.Mongo.Collection("applet")
}

//保存应用
func (applet *Applet) Save() error {
	timestamp := time.Now().Unix()
	applet.CreateIn = timestamp
	applet.UpdateIn = timestamp
	applet.Version = 1
	coll := applet.collection()
	if _, err := coll.InsertOne(context.Background(), applet); err != nil {
		return err
	}
	return nil
}
