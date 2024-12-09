package httpx

import (
	"errors"
	"strings"

	enLang "github.com/go-playground/locales/en"
	zhLang "github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
)

var supportLang map[string]string

// 自定义验证器
type CustomValidator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     map[string]ut.Translator
	ErrorCode int
}

func NewValidator() *CustomValidator {
	v := CustomValidator{}
	// 设置默认错误代码
	v.ErrorCode = 1001

	en := enLang.New()
	zh := zhLang.New()
	v.Uni = ut.New(zh, en, zh)
	v.Validator = validator.New()
	enTrans, _ := v.Uni.GetTranslator("en")
	zhTrans, _ := v.Uni.GetTranslator("zh")
	v.Trans = make(map[string]ut.Translator)
	v.Trans["en"] = enTrans
	v.Trans["zh"] = zhTrans
	// 添加语言支持
	initSupportLanguages()

	err := enTranslations.RegisterDefaultTranslations(v.Validator, enTrans)
	if err != nil {
		logx.Errorw("register English translation failed", logx.Field("detail", err.Error()))
		return nil
	}
	err = zhTranslations.RegisterDefaultTranslations(v.Validator, zhTrans)
	if err != nil {
		logx.Errorw("register Chinese translation failed", logx.Field("detail", err.Error()))

		return nil
	}

	return &v
}

func (v *CustomValidator) Validate(data any, lang string) string {
	err := v.Validator.Struct(data)
	if err == nil {
		return ""
	}

	targetLang, parseErr := ParseAcceptLanguage(lang)
	if parseErr != nil {
		return parseErr.Error()
	}

	errs, ok := err.(validator.ValidationErrors)

	if ok {
		transData := errs.Translate(v.Trans[targetLang])
		s := strings.Builder{}
		for _, v := range transData {
			s.WriteString(v)
			s.WriteString(" ")
		}
		return s.String()
	}

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return invalid.Error()
	}

	return ""
}

func ParseAcceptLanguage(lang string) (string, error) {
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		return "", errors.New("fail to parse accept language")
	}

	for _, v := range tags {
		if val, ok := supportLang[v.String()]; ok {
			return val, nil
		}
	}

	return "zh", nil
}

func initSupportLanguages() {
	supportLang = make(map[string]string)
	supportLang["zh"] = "zh"
	supportLang["zh-CN"] = "zh"
	supportLang["en"] = "en"
	supportLang["en-US"] = "en"
}

// 将验证函数注册到验证器
func RegisterValidation(tag string, fn validator.Func) {
	if err := xValidator.Validator.RegisterValidation(tag, fn); err != nil {
		logx.Must(errors.Join(err, errors.New("failed to register the validation function, tag is "+tag)))
	}
}

// 将验证翻译注册到验证器
func RegisterValidationTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc,
	translationFn validator.TranslationFunc) {
	if err := xValidator.Validator.RegisterTranslation(tag, trans, registerFn, translationFn); err != nil {
		logx.Must(errors.Join(err, errors.New("failed to register the validation translation, tag is "+tag)))
	}
}

// 设置发生错误时验证器的错误代码
func SetValidatorErrorCode(code int) {
	xValidator.ErrorCode = code
}
