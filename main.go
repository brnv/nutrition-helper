package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/brnv/nutrition-helper/nutrition"
	"github.com/gocraft/web"
)

type Context struct{}

func handleRoot(w web.ResponseWriter, r *web.Request) {
	handleMeal(w, r)
}

func handleMeal(w web.ResponseWriter, r *web.Request) {
	r.Header.Set("Content-Type", "text/html")
	templateIndex().Execute(w, templateContent("meal"))
}

func handleStats(w web.ResponseWriter, r *web.Request) {
	r.Header.Set("Content-Type", "text/html")
	s := &nutrition.StorageRecord{Id: 1}
	nutrition.Get(s)

	templateIndex().Execute(w, templateContent("settings"))
}

func handleSettings(w web.ResponseWriter, r *web.Request) {
	r.Header.Set("Content-Type", "text/html")
	templateIndex().Execute(w, templateContent("settings"))
}

func handleEat(w web.ResponseWriter, r *web.Request) {
	nutrition.Eat(r)
	http.Redirect(w, r.Request, "/", http.StatusFound)
}

func main() {
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddleware("www")).
		Get("/", handleRoot).
		Get("/meal", handleMeal).
		Get("/stats", handleStats).
		Get("/settings", handleSettings).
		Post("/eat", handleEat)

	log.Fatal(http.ListenAndServe(":8084", router))
}

func templateIndex() *template.Template {
	f, _ := os.Open("www/templates/index.tpl")
	raw, _ := ioutil.ReadAll(f)
	return template.Must(template.New("index").Parse(string(raw)))
}

func templateContent(template string) interface{} {
	f, _ := os.Open("www/templates/" + template + ".tpl")
	content, _ := ioutil.ReadAll(f)

	f, _ = os.Open("www/templates/menu.tpl")
	menu, _ := ioutil.ReadAll(f)

	return struct {
		Content string
		Menu    string
	}{
		Content: string(content),
		Menu:    string(menu),
	}
}
