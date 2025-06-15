package dom

type (
	MediaDevices struct {
		DOM
	}
)

func GetMediaDevices() MediaDevices {

	return GetNavigator().MediaDevices()
}
