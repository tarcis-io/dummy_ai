package util

func IsDesktop() bool {

	return false
}

func IsMobile() bool {

	return !IsDesktop()
}
