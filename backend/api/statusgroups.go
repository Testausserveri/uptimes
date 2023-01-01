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

func (s *Server) provideRoutes(c echo.Context) error {
	c.JSON(http.StatusOK, map[string]any{
		"providedRoutes": s.handledGroups,
	})
	return nil
}

func (s *Server) HandleStatusGroup(sg *APIStatusGroup) {
	if sg == nil {
		log.Fatal("adding nil handler is not allowed")
	}

	for _, servepath := range s.handledGroups {
		if servepath == sg.ServePath {
			log.Println("cannot initialize two statusgroups with same serve path")
			return
		}
	}

	s.handledGroups = append(s.handledGroups, sg.ServePath)
	s.router.GET(sg.ServePath, sg.Handler)
}
