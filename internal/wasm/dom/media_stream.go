package dom

type (
	MediaStream struct {
		DOM
	}
)

func GetUserMedia(constraints map[string]any) (MediaStream, error) {

	return GetMediaDevices().GetUserMedia(constraints)
}
