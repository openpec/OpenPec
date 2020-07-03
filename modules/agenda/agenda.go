package agenda

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/OpenPEC/config"
)

//HomeGet é o handler da pagina principal
func HomeGet(srv *config.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Pega a pasta raiz do projeto
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		//template html
		t, err := template.ParseFiles(wd + "/templates/agenda/agenda.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	}
}
