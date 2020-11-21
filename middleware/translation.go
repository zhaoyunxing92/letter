package middleware

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/zhaoyunxing92/letter/global"
)

import (
	"github.com/gin-gonic/gin"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

/*
多语言翻译器
*/
func TranslationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//设置国际化翻译器
		uni := ut.New(en.New(), zh.New())
		val := validator.New()

		//获取语言，默认中文
		locale := ctx.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "zh":
			zh_trans.RegisterDefaultTranslations(val, trans)
			break
		case "en":
			en_trans.RegisterDefaultTranslations(val, trans)
			break
		default:
			zh_trans.RegisterDefaultTranslations(val, trans)
		}
		ctx.Set(global.TranslatorKey, trans)
		ctx.Set(global.ValidatorKey, val)
		ctx.Next()
	}
}
