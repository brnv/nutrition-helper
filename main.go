package main

import (
	"log"
	"net/http"

	"github.com/brnv/nutrition-helper/nutrition"
	"github.com/gocraft/web"
)

type Context struct {
}

func handlePost(w web.ResponseWriter, r *web.Request) {
	r.ParseForm()
	nutrition.Eat(r.Form)
	http.Redirect(w, r.Request, "/", http.StatusFound)
}

func main() {
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddleware("www")).
		Post("/eat", handlePost)

	log.Fatal(http.ListenAndServe(":8080", router))
}
