package trans

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
)

func ZhTranslator() ut.Translator {
	//enUs := en_US.New()
	zhCn := zh_Hans_CN.New()
	translator := ut.New(zhCn, zhCn)
	zhTranlator, found := translator.GetTranslator("zh_Hans_CN")
	if found == true {
		return translator.GetFallback()
	}

	return zhTranlator
}
