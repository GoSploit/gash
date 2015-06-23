// gosh project main.go
package main

import (
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

	theme, _ = tartheme.LoadDir("themes/default")

	static := theme.Prefix("static")

	templates = theme.Prefix("templates").Templates()

	router.HandleFunc("/", Index)
	router.Handle("/static/", http.StripPrefix("/static/", static))

	router.HandleFunc("/api/GetWork", GetWork)
}

func Index(rw http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(rw, "index.tpl", nil)
}
