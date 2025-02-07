package util

const (
	app = iota
	appDescription
)

var (
	allMessages = map[string]map[int]string{
		English:    {},
		Spanish:    {},
		Portuguese: {},
	}
	messages = allMessages[language]
)

func App() string {

	return messages[app]
}

func AppDescription() string {

	return messages[appDescription]
}
