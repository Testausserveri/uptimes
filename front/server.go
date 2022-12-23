package front

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Testausserveri/uptimes/engine"
	"github.com/Testausserveri/uptimes/storage"
)

var usedPaths []string = []string{"/assets/"}

func newHandler(group engine.StatusGroup) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templatePath := fmt.Sprintf("public/%s", group.Config.TemplateName)
		t, err := template.ParseFiles(templatePath)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := t.Execute(w, struct {
			Storage storage.Storage
		}{
			*group.Storage,
		}); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func InitRoute(statusGroup engine.StatusGroup) {
	for _, rn := range usedPaths {
		if rn == statusGroup.Config.ServePath {
			log.Fatalln(rn, "is already being served.")
		}
	}

	http.Handle(statusGroup.Config.ServePath, newHandler(statusGroup))

	usedPaths = append(usedPaths, statusGroup.Config.ServePath)
}

func Serve(address string, port int, eng ...engine.StatusGroup) error {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets/"))))

	return http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil)
}
