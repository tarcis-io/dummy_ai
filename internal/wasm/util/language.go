package util

import (
	"syscall/js"
)

const (
	English    = "en"
	Spanish    = "es"
	Portuguese = "pt"
)

const (
	DefaultLanguage = English
)

var (
	language = DefaultLanguage
)

func Language() string {

	return language
}

func SetLanguage(language string) {

	js.Global().Get("localStorage").Call("setItem", "language", language)
}
