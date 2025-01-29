package util

const (
	app = iota
	appDescription
	appNavigation
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:            "DummyAI",
			appDescription: "Artificial intelligence for dummies",
			appNavigation:  "Navigation",
		},
		Spanish: {
			app:            "DummyAI",
			appDescription: "Artificial intelligence for dummies",
			appNavigation:  "Navegación",
		},
		Portuguese: {
			app:            "DummyAI",
			appDescription: "Artificial intelligence for dummies",
			appNavigation:  "Navegação",
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

func AppNavigation() string {

	return messages[appNavigation]
}
