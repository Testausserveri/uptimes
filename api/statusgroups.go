package api

import (
	"log"
	"net/http"

	"github.com/Testausserveri/uptimes/types"
	"github.com/labstack/echo"
)

type APIStatusGroup struct {
	ServePath string
	types.StatusGroup
}

func NewAPIStatusGroup(sg *types.StatusGroup, p string) *APIStatusGroup {
	return &APIStatusGroup{
		ServePath:   p,
		StatusGroup: *sg,
	}
}

func (s *APIStatusGroup) Handler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, s.StatusGroup)
}

func (s *Server) HandleStatusGroup(sg *APIStatusGroup) {
	if sg == nil {
		log.Fatal("adding nil handler is not allowed")
	}

	s.router.GET(sg.ServePath, sg.Handler)
}
