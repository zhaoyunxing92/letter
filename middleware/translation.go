package middleware

import "github.com/zhaoyunxing92/letter/public"

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/zh"
)

/*
多语言翻译器
*/
func TranslationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale := ctx.DefaultQuery("locale", "zh")

		trans, _ := public.Uni.GetTranslator(locale)

		switch locale {
		case "zh":
			zh.RegisterDefaultTranslations(public.Validate, trans)
			break
		case "en":
			en.RegisterDefaultTranslations(public.Validate, trans)
			break
		default:
			zh.RegisterDefaultTranslations(public.Validate, trans)
		}
		ctx.Set("trans",trans)
		ctx.Next()
	}
}
