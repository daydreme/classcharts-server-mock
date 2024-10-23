package frontend

import (
	"github.com/CommunityCharts/CCModels/shared"
	"html/template"
	"net/http"
	"os"
	"path"
)

var cwd, _ = os.Getwd()

//var templates = template.Must(template.ParseFiles(file("index")))

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "index", &shared.Object{})
}

func AnnouncementsPageHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "announce", &shared.Object{})
}

func render(w http.ResponseWriter, name string, data *shared.Object) {
	t, err := template.ParseFiles(file(name), file("head"))
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func file(name string) string {
	return path.Join(cwd, "pkg", "frontend", "views", name+".html")
}

func Dir(name string) string {
	return path.Join(cwd, "pkg", "frontend", name)
}
