package param

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/letter/global"
)

//创建应用参数
type NewApplet struct {
	CorpId    string `json:"corpId" form:"corpId" validate:"required"`       //企业id
	AgentId   string `json:"agentId" form:"agentId" validate:"required"`     //应用id
	AppKey    string `json:"appKey" form:"appKey" validate:"required"`       //应用key
	AppSecret string `json:"appSecret" form:"appSecret" validate:"required"` //应用秘钥
}

//更新应用
type UpdateApplet struct {
	CorpId    string `json:"corpId" form:"corpId" omitempty:"required"`      //企业id
	AgentId   string `json:"agentId" form:"agentId" omitempty:"required"`    //应用id
	AppSecret string `json:"appSecret" form:"appSecret" validate:"required"` //应用秘钥
}

//参数验证
func (args *NewApplet) Validate(ctx *gin.Context) error {
	return global.DefaultGetValidParams(ctx, args)
}

// 更新参数验证
func (args *UpdateApplet) Validate(ctx *gin.Context) error {
	return global.DefaultGetValidParams(ctx, args)
}
