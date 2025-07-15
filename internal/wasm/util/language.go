package util

import (
	"strings"

	"dummy_ai/internal/env"
	"dummy_ai/internal/wasm/dom"
)

type (
	Language struct {
		code string
		name string
	}
)

func (language Language) Code() string {
	return language.code
}

func (language Language) Name() string {
	return language.name
}

var (
	english = Language{
		code: "en",
		name: "English",
	}
	spanish = Language{
		code: "es",
		name: "Español",
	}
	portuguese = Language{
		code: "pt",
		name: "Português",
	}
)

var (
	languages = []Language{
		english,
		spanish,
		portuguese,
	}
)

func English() Language {
	return english
}

func Spanish() Language {
	return spanish
}

func Portuguese() Language {
	return portuguese
}

func Languages() []Language {
	return languages
}

func GetLanguage() Language {
	if language, found := lookupStorageLanguage(); found {
		return language
	}
	if language, found := lookupNavigatorLanguage(); found {
		return language
	}
	if language, found := lookupEnvironmentLanguage(); found {
		return language
	}
	return english
}

func SetLanguage(language Language) {
	dom.GetLocalStorage().SetItem("language", language.code)
}

func LookupLanguage(code string) (Language, bool) {
	for _, language := range languages {
		if language.code == code {
			return language, true
		}
	}
	return Language{}, false
}

func lookupStorageLanguage() (Language, bool) {
	if language, found := dom.GetLocalStorage().GetItem("language"); found {
		return LookupLanguage(language)
	}
	return Language{}, false
}

func lookupNavigatorLanguage() (Language, bool) {
	if language, found := dom.GetNavigator().Language(); found {
		return LookupLanguage(strings.Split(language, "-")[0])
	}
	return Language{}, false
}

func lookupEnvironmentLanguage() (Language, bool) {
	return LookupLanguage(env.Language())
}
