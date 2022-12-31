package api

import (
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
)

type Server struct {
	dbConnection  *pgx.Conn
	listenAddress string
	router        *echo.Echo
}

func NewServer(dbConn *pgx.Conn, listenAddress string, router *echo.Echo) *Server {
	return &Server{
		dbConnection:  dbConn,
		listenAddress: listenAddress,
		router:        router,
	}
}

func (s *Server) Start() error {
	return s.router.Start(s.listenAddress)
}

func (s *Server) AddRoute(m, p string, r echo.HandlerFunc) {
	s.router.Add(m, p, r)
}
