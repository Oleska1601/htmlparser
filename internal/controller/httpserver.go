package controller

import (
	"fmt"
	"htmlparser/internal/logging"
	"net/http"
)

type Server struct {
	Server *http.Server
	u      UsecaseInterface
	l      logging.LoggerInterface
}

func New(host string, port int, u UsecaseInterface, l logging.LoggerInterface) *Server {
	s := &Server{
		Server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
		},
		u: u,
		l: l,
	}
	s.setUpRouter()
	return s

}
