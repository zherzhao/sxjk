package trans

import (
	"webconsole/global"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"

	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func Init() error {
	var err error
	uni := ut.New(en.New(), zh.New())
	locale := global.ServerSetting.Locale
	trans, _ := uni.GetTranslator(locale)

	// 修改gin框架中的Validator引擎属性，实现自定制
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
	}
	// 注册翻译器
	global.Trans = trans
	return err

}
