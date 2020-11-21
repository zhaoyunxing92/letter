package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoyunxing92/letter/domain"
	"github.com/zhaoyunxing92/letter/global"
	"github.com/zhaoyunxing92/letter/param"
)

//创建应用
func CreateApplet(ctx *gin.Context) {
	var args param.NewApplet
	//参数验证
	if err := args.Validate(ctx); err != nil {
		global.ResponseError(ctx, 400, err)
		ctx.Abort()
		return
	}
	//数据保存
	applet := new(domain.Applet)
	applet.AgentId = args.AgentId
	applet.AppSecret = args.AppSecret
	applet.AppKey = args.AppKey
	applet.CorpId = args.CorpId
	if err := applet.Save(); err != nil {
		global.ResponseError(ctx, 400, err)
		ctx.Abort()
		return
	}
	global.ResponseSuccess(ctx, applet)
}
