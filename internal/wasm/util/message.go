package util

const (
	app = iota
	appDescription
	navigationItemAbout
	navigationItemHome
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			navigationItemAbout: "About",
			navigationItemHome:  "Home",
		},
		Spanish: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			navigationItemAbout: "About",
			navigationItemHome:  "Home",
		},
		Portuguese: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			navigationItemAbout: "Sobre",
			navigationItemHome:  "Início",
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
