package util

const (
	app = iota
	appDescription
	appButtonNavigationAriaLabel
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                          "DummyAI",
			appDescription:               "Artificial intelligence for dummies",
			appButtonNavigationAriaLabel: "Navigation",
		},
		Spanish: {
			app:                          "DummyAI",
			appDescription:               "Artificial intelligence for dummies",
			appButtonNavigationAriaLabel: "Navegación",
		},
		Portuguese: {
			app:                          "DummyAI",
			appDescription:               "Artificial intelligence for dummies",
			appButtonNavigationAriaLabel: "Navegação",
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

func AppButtonNavigationAriaLabel() string {

	return messages[appButtonNavigationAriaLabel]
}
