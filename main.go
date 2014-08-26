package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brnv/nutrition-helper/calculator"
	"github.com/gocraft/web"
)

type Context struct {
}

func handlePost(w web.ResponseWriter, r *web.Request) {
	r.ParseForm()
	var meal struct {
		name  string
		facts calculator.NutritionFacts
	}
	meal.name = r.FormValue("productInput")
	fmt.Fprint(w, meal)
}

func main() {
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddleware("www")).
		Post("/eat", handlePost)

	log.Fatal(http.ListenAndServe(":8080", router))
}
