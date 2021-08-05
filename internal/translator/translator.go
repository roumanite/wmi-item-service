package translator

import (
	"bytes"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
	idt "github.com/go-playground/validator/v10/translations/id"
	"gopkg.in/yaml.v3"
)

var (
	enTrans            = en.New()
	uni                = ut.New(enTrans, enTrans)
	enTranslator, _    = uni.GetTranslator("en")
	translators        = map[string]ut.Translator{"en": enTranslator}
	loadedTranslations = make(map[string]map[string]string)
)

func LoadTranslations(commonPath string, validationPath string) error {
	files, err := ioutil.ReadDir(validationPath)
	if err != nil {
		return err
	}

	for _, filePath := range files {
		lang := strings.TrimSuffix(filePath.Name(), ".yaml")
		f, err := ioutil.ReadFile(validationPath + "/" + filePath.Name())
		if err != nil {
			return err
		}
		var trans map[string]string
		err = yaml.Unmarshal(f, &trans)
		if err != nil {
			return err
		}
		loadedTranslations[lang] = trans
		translators[lang] = getTranslator(lang)
	}

	err = uni.Import(ut.FormatJSON, commonPath)
	if err != nil {
		return err
	}

	err = uni.VerifyTranslations()
	if err != nil {
		return err
	}

	return nil
}

func RegisterTranslations(v *validator.Validate) {
	for lang, translations := range loadedTranslations {
		registerTranslations(lang, v)

		for key, translation := range translations {
			v.RegisterTranslation(key, translators[lang], func(ut ut.Translator) error {
				return ut.Add(key, translation, true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				fieldName := toHumanReadable(fe.Field())
				translatedFName, err := ut.T(fe.Field())
				if err == nil {
					fieldName = capitalizeFirst(translatedFName)
				}
				if len(fe.Param()) == 0 {
					t, _ := ut.T(fe.Tag(), fieldName)
					return t
				}

				paramName := lowerFirst(fe.Param())
				translatedParamName, err := ut.T(paramName)
				if err == nil {
					paramName = translatedParamName
				}
				t, _ := ut.T(fe.Tag(), fieldName, paramName)
				return t
			})
		}
	}
}

func Translate(fe validator.FieldError, lang string) string {
	translator, ok := translators[lang]
	if !ok {
		return fe.Translate(translators["en"])
	}
	return fe.Translate(translator)
}

func registerTranslations(lang string, v *validator.Validate) {
	translator, ok := translators[lang]
	if !ok {
		ent.RegisterDefaultTranslations(v, translators["en"])
	}

	switch lang {
	case "en":
		ent.RegisterDefaultTranslations(v, translator)
	case "id":
		idt.RegisterDefaultTranslations(v, translator)
	default:
		ent.RegisterDefaultTranslations(v, translator)
	}
}

func getTranslator(lang string) ut.Translator {
	if lang == "en" {
		return translators["en"]
	}

	if lang == "id" {
		idTrans := id.New()
		uni.AddTranslator(idTrans, true)
		idTranslator, _ := uni.GetTranslator("id")
		return idTranslator
	}

	return translators["en"]
}

func capitalizeFirst(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func toHumanReadable(s string) string {
	buf := &bytes.Buffer{}
	for i, rune := range s {
		if i == 0 {
			buf.WriteRune(unicode.ToUpper(rune))
			continue
		}

		if unicode.IsUpper(rune) {
			buf.WriteRune(' ')
		}
		buf.WriteRune(unicode.ToLower(rune))
	}
	return buf.String()
}

func lowerFirst(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToLower(tmp[0])
	return string(tmp)
}