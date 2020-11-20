package domain

import (
	"github.com/zhaoyunxing92/letter/public"
	"go.mongodb.org/mongo-driver/mongo"
)

type Applet struct {
	Name    string
	CorpId  string
	AgentId string
}

func (a *Applet) GetCollection() *mongo.Collection  {
	return public.Mongo.Collection("applet")
}
