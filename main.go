package main

import "net/http"

func main() {

	http.Handle("/", http.FileServer(http.Dir("./www")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
