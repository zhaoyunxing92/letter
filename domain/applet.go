package domain

import (
	"context"
	"github.com/zhaoyunxing92/letter/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Applet struct {
	base
	CorpId         string `json:"corpId"`            //企业id
	AgentId        uint64 `json:"agentId"`           //应用id
	AppKey         string `json:"appKey" bson:"_id"` //应用key
	AppSecret      string `json:"appSecret"`         //应用秘钥
	Name           string `json:"name"`              //应用名称
	Icon           string `json:"icon"`              //应用图标
	Desc           string `json:"desc"`              //应用描述
	Self           bool   `json:"self"`              //是否是自建应用false：不是,true：自建
	OmpLink        string `json:"ompLink"`           //应用的OA后台管理主页
	HomepageLink   string `json:"homepageLink"`      //应用的移动端主页。
	PcHomepageLink string `json:"pcHomepageLink"`    //应用的PC端主页。

}

//获取集合
func (applet *Applet) collection() *mongo.Collection {
	return global.Mongo.Collection("applet")
}

//保存应用
func (applet *Applet) Save() error {
	coll := applet.collection()
	applet.CreateIn = time.Now().Unix()
	applet.Version = 1
	if _, err := coll.InsertOne(context.Background(), applet); err != nil {
		return err
	}
	return nil
}

//根据id查询
func (applet *Applet) FindById() error {
	coll := applet.collection()

	if err := coll.FindOne(context.Background(), bson.M{"_id": applet.AppKey}).Decode(&applet); err != nil {
		return err
	}
	return nil
}
