package ctrl

import (
	"errors"
	"fmt"
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

//根据id获取应用
func GetApplet(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		global.ResponseError(ctx, 400, errors.New("id不能为空"))
		ctx.Abort()
		return
	}

	applet := new(domain.Applet)
	applet.AppKey = id

	if err := applet.FindById(); err != nil {
		if "mongo: no documents in result" == err.Error() {
			err = errors.New(fmt.Sprintf("id:[%s]不存在", id))
		}
		global.ResponseError(ctx, 400, err)
		ctx.Abort()
		return
	}
	global.ResponseSuccess(ctx, applet)
}

//根据id获取应用
func UpdateApplet(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		global.ResponseError(ctx, 400, errors.New("id不能为空"))
		ctx.Abort()
		return
	}
	
	//参数验证
	var args param.UpdateApplet
	if err := args.Validate(ctx); err != nil {
		global.ResponseError(ctx, 400, err)
		ctx.Abort()
		return
	}

	applet := new(domain.Applet)
	applet.AppKey = id

	if err := applet.FindById(); err != nil {
		if "mongo: no documents in result" == err.Error() {
			err = errors.New(fmt.Sprintf("id:[%s]不存在", id))
		}
		global.ResponseError(ctx, 400, err)
		ctx.Abort()
		return
	}
	global.ResponseSuccess(ctx, applet)
}
