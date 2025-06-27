package server

type (
	Server struct{}
)

func (server Server) ListenAndServe() {
}

func New() Server {

	return Server{}
}
