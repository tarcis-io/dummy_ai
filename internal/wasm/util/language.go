package util

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
