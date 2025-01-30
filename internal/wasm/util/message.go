package util

const (
	app = iota
	appDescription
	mastheadNavigationButtonAriaLabel
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navigation",
		},
		Spanish: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navegación",
		},
		Portuguese: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navegação",
		},
	}
	messages = allMessages[language]
)

func App() string {

	return messages[app]
}

func AppDescription() string {

	return messages[appDescription]
}

func MastheadNavigationButtonAriaLabel() string {

	return messages[mastheadNavigationButtonAriaLabel]
}
