package public

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func init() {
	log.Println("init validate")
	en := en.New()
	zh := zh.New()
	Uni = ut.New(en, zh)
	Validate = validator.New()
}
