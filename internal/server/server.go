package server

type (
	Server struct {
	}
)

func New() (*Server, error) {
	server := &Server{}
	return server, nil
}
