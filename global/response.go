package global

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Code int

const (
	SuccessCode Code = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvalidRequestErrorCode = 401
	CustomizeCode           = 1000
	GROUPALL_SAVE_FLOWERROR = 2001
)

//统一返回
type Response struct {
	Code    Code        `json:"code"`
	Msg     string      `json:"msg"`
	TraceId string      `json:"tranceId"`
	Data    interface{} `json:"data"`
	Time    string      `json:"time"`
}

func ResponseError(ctx *gin.Context, code Code, err error) {
	tranceId := ctx.GetString(TranceIdKey)

	resp := &Response{code, err.Error(), "", tranceId, time.Now().Format("2006-01-02 15:04:05")}
	ctx.JSON(200, resp)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	tranceId := ctx.GetString(TranceIdKey)

	resp := &Response{SuccessCode, "ok", tranceId, data, time.Now().Format("2006-01-02 15:04:05")}
	ctx.JSON(200, resp)
}
