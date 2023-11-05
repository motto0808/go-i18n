package i18n

import (
	"embed"
	"fmt"
	"log"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/ar"
	"github.com/go-playground/locales/currency"
	"github.com/go-playground/locales/de"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/ko"
	"github.com/go-playground/locales/ms"
	"github.com/go-playground/locales/pt"
	"github.com/go-playground/locales/ru"
	"github.com/go-playground/locales/th"
	"github.com/go-playground/locales/vi"
	"github.com/go-playground/locales/zh_Hans"
	"github.com/go-playground/locales/zh_Hant"
	ut "github.com/go-playground/universal-translator"
)

var (
	utrans *ut.UniversalTranslator
)

var localeMap = map[string]func() locales.Translator{
	"en":      en.New,
	"zh_Hans": zh_Hans.New,
	"zh_Hant": zh_Hant.New,
	"ja":      ja.New,
	"es":      es.New,
	"pt":      pt.New,
	"fr":      fr.New,
	"de":      de.New,
	"ar":      ar.New,
	"ru":      ru.New,
	"ko":      ko.New,
	"ms":      ms.New,
	"th":      th.New,
	"id":      id.New,
	"vi":      vi.New,
}

//go:embed assets/*
var assets embed.FS

// 加载所有语言
func InitAll(defaultLang string) {
	def := localeMap[defaultLang]()
	supportedLocales := []locales.Translator{}
	for _, v := range localeMap {
		supportedLocales = append(supportedLocales, v())
	}
	utrans = ut.New(def, supportedLocales...)

	dirs, err := assets.ReadDir("assets")
	if err != nil {
		log.Fatalf("i18n import translations error: %s", err)
		return
	}
	for _, dir := range dirs {
		filename := fmt.Sprintf("assets/%s/time.json", dir.Name())
		f, err := assets.Open(filename)
		if err != nil {
			continue
		}
		err = utrans.ImportByReader(ut.FormatJSON, f)
		if err != nil {
			log.Panicf("i18n import translations %s, %v", filename, err)
		}
	}

	verify()
}

// 加载指定语言
// 传入的第一个参数会作为fallback
func Init(langs ...string) {
	if len(langs) == 0 {
		return
	}
	def := localeMap[langs[0]]()
	supportedLocales := []locales.Translator{}
	for _, v := range langs {
		supportedLocales = append(supportedLocales, localeMap[v]())
	}
	utrans = ut.New(def, supportedLocales...)

	for _, v := range langs {
		filename := fmt.Sprintf("assets/%s/time.json", v)
		f, err := assets.Open(filename)
		if err != nil {
			continue
		}
		err = utrans.ImportByReader(ut.FormatJSON, f)
		if err != nil {
			log.Panicf("i18n import translations %s %v", filename, err)
		}
	}

	verify()
}

type RelativeTimeFormatUnit = int

const (
	RelativeTimeSecond RelativeTimeFormatUnit = iota
	RelativeTimeMinute
	RelativeTimeHour
	RelativeTimeDay
	RelativeTimeWeek
)

// Translator wraps ut.Translator in order to handle errors transparently
// it is totally optional but recommended as it can now be used directly in
// templates and nobody can add translations where they're not supposed to.
type Translator interface {
	locales.Translator

	// creates the translation for the locale given the 'key' and params passed in.
	// wraps ut.Translator.T to handle errors
	T(key interface{}, params ...string) string

	// creates the cardinal translation for the locale given the 'key', 'num' and 'digit' arguments
	//  and param passed in.
	// wraps ut.Translator.C to handle errors
	C(key interface{}, num float64, digits uint64, param string) string

	// creates the ordinal translation for the locale given the 'key', 'num' and 'digit' arguments
	// and param passed in.
	// wraps ut.Translator.O to handle errors
	O(key interface{}, num float64, digits uint64, param string) string

	//  creates the range translation for the locale given the 'key', 'num1', 'digit1', 'num2' and
	//  'digit2' arguments and 'param1' and 'param2' passed in
	// wraps ut.Translator.R to handle errors
	R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string

	// Currency returns the type used by the given locale.
	Currency() currency.Type

	// 格式化一段时间
	FmtDuration(seconds int64) string

	// 格式化指定的相对时间
	FmtRelativeTime(value int64, unit RelativeTimeFormatUnit) string
}

// implements Translator interface definition above.
type translator struct {
	locales.Translator
	trans ut.Translator
}

var _ Translator = (*translator)(nil)

func (t *translator) T(key interface{}, params ...string) string {

	s, _ := t.trans.T(key, params...)

	return s
}

func (t *translator) C(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.Printf("issue translating cardinal key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) O(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.Printf("issue translating ordinal key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string {

	s, err := t.trans.R(key, num1, digits1, num2, digits2, param1, param2)
	if err != nil {
		log.Printf("issue translating range key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) Currency() currency.Type {

	// choose your own locale. The reason it isn't mapped for you is because many
	// countries have multiple currencies; it's up to you and you're application how
	// and which currencies to use. I recommend adding a function it to to your custon translator
	// interface like defined above.
	switch t.Locale() {
	case "zh_Hans":
		return currency.CNY
	case "zh_Hant":
		return currency.HKD
	case "en":
		return currency.USD
	case "ja":
		return currency.JPY
	case "es", "fr", "pt", "de":
		return currency.EUR
	case "th":
		return currency.THB
	case "ko":
		return currency.KRW
	case "ms":
		return currency.MYR
	case "ar":
		return currency.AED
	case "ru":
		return currency.RUB
	default:
		return currency.USD
	}
}

func (t *translator) FmtDuration(seconds int64) string {
	var ft func(int64) string
	ft = func(second int64) string {
		switch {
		case second <= 0:
			return ""
		case second < 60:
			return t.C("seconds-num", float64(second), 0, t.FmtNumber(float64(second), 0))
		case second < 3600:
			n := float64(second / 60)
			return t.C("minutes-num", n, 0, t.FmtNumber(n, 0))
		case second < 86400:
			n := float64(second / 60 / 60)
			fmtStr := t.fmtDurationStr()
			return fmt.Sprintf(fmtStr, t.C("hours-num", n, 0, t.FmtNumber(n, 0)), ft(second%3600))
		default:
			n := float64(second / 60 / 60 / 24)
			fmtStr := t.fmtDurationStr()
			return fmt.Sprintf(fmtStr, t.C("days-num", n, 0, t.FmtNumber(n, 0)), ft(second%(3600*24)))
		}
	}
	return ft(seconds)
}

func (t *translator) fmtDurationStr() string {
	if t.Locale() == "zh_Hans" || t.Locale() == "zh_Hant" || t.Locale() == "ja" || t.Locale() == "ko" {
		return "%s%s"
	}
	return "%s %s"
}

func verify() {
	err := utrans.VerifyTranslations()
	if err != nil {
		log.Fatalf("i18n verify translations %v", err)
	}
}

func GetTranslator(lang ...string) Translator {
	tr, ok := utrans.FindTranslator(lang...)
	if !ok {
		return nil
	}
	trans := translator{trans: tr, Translator: tr.(locales.Translator)}
	return &trans
}

func AvaliableLocales() []string {
	keys := make([]string, 0, len(localeMap))
	for k := range localeMap {
		keys = append(keys, k)
	}
	return keys
}
