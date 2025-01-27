package util

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
