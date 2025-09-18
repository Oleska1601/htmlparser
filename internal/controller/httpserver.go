package controller

import (
	"fmt"
	"htmlparser/pkg/logger"
	"net/http"
)

type Server struct {
	Server *http.Server
	u      UsecaseInterface
	l      logger.LoggerInterface
}

func New(host string, port int, u UsecaseInterface, l logger.LoggerInterface) *Server {
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
