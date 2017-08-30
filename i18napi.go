package i18napi

import (
	"github.com/nicksnyder/go-i18n/i18n"
)

var tmap = make(map[string]i18n.TranslateFunc)
var defaultfunc i18n.TranslateFunc = emptyTran

func emptyTran(translationID string, args ...interface{}) string {
	return translationID
}

// SetDefalut 设置默认语言
func SetDefalut(lang string) {
	tfunc, ok := tmap[lang]
	if ok == true {
		defaultfunc = tfunc
	}
}

// AddTranLang 添加翻译语言
func AddTranLang(lang string, fpath string) {

	i18n.MustLoadTranslationFile(fpath)
	tfunc, _ := i18n.Tfunc(lang)
	tmap[lang] = tfunc
}

// TranID 根据ID转换
func TranID(lang string, id string) string {
	tfunc, ok := tmap[lang]
	if ok == true {
		return tfunc(id)
	}
	return defaultfunc(id)
}

// TranOne 转换
func TranOne(lang string, id string, i interface{}) string {
	tfunc, ok := tmap[lang]
	if ok == true {
		return tfunc(id, i)
	}
	return defaultfunc(id)
}

// TranMap 根据map对应情况进行转换
func TranMap(lang string, id string, m map[string]interface{}) string {

	tfunc, ok := tmap[lang]
	if ok == true {
		return tfunc(id, m)
	}
	return defaultfunc(id)
}
