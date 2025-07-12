package dom

type (
	Navigator struct {
		DOM
	}
)

func (navigator Navigator) Language() (string, bool) {
	if language := navigator.Get("language"); language.Truthy() {
		return language.String(), true
	}
	return "", false
}

func GetNavigator() Navigator {
	return GetWindow().Navigator()
}
