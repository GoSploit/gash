// gosh project main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/andyleap/tartheme"
	"github.com/gorilla/mux"
)

var (
	theme     *tartheme.TarTheme
	templates *template.Template

	Projects []*Project
	Workers  *WorkerMap
)

func main() {
	router := mux.NewRouter()
	Workers = &WorkerMap{
		workermap: make(map[string]*Worker),
	}

	theme, _ = tartheme.LoadDir("themes/default")

	static := theme.Prefix("static/")

	templates = theme.Prefix("templates/").Templates()

	router.HandleFunc("/", Index)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	router.HandleFunc("/agent/register", Register)
	router.HandleFunc("/agent/getwork", GetWork)
	router.HandleFunc("/agent/finishwork", FinishWork)

	http.ListenAndServe(":5126", router)
}

func Index(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(templates.ExecuteTemplate(rw, "index.tpl", nil))
}

func NewProject(rw http.ResponseWriter, req *http.Request) {

}
