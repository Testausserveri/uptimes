package api

import (
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
)

type Server struct {
	listenAddress string

	handledGroups []string
	dbConnection  *pgx.Conn
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
	s.router.GET("/routes", s.provideRoutes)
	return s.router.Start(s.listenAddress)
}
