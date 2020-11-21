package global

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"strings"
)

//参数验证
func DefaultGetValidParams(ctx *gin.Context, params interface{}) error {
	if err := ctx.ShouldBind(params); err != nil {
		return err
	}
	//获取验证器
	valid, err := getValidator(ctx)
	if err != nil {
		return err
	}
	//获取翻译器
	trans, err := getTranslation(ctx)
	if err != nil {
		return err
	}
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

// 获取验证器
func getValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok := c.Get(ValidatorKey)
	if !ok {
		return nil, errors.New("未设置验证器")
	}
	validate, ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("获取验证器失败")
	}
	return validate, nil
}

// 获取翻译器
func getTranslation(c *gin.Context) (ut.Translator, error) {
	trans, ok := c.Get(TranslatorKey)
	if !ok {
		return nil, errors.New("未设置翻译器")
	}
	translator, ok := trans.(ut.Translator)
	if !ok {
		return nil, errors.New("获取翻译器失败")
	}
	return translator, nil
}
