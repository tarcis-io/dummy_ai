package util

const (
	app = iota
	appDescription
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:            "DummyAI",
			appDescription: "Artificial intelligence for dummies",
		},
		Spanish: {
			app:            "DummyAI",
			appDescription: "Artificial intelligence for dummies",
		},
		Portuguese: {
			app:            "DummyAI",
			appDescription: "Artificial intelligence for dummies",
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
