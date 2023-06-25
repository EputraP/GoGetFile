package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	fs := http.FileServer(http.Dir("./image"))
	http.Handle("/", cors(fs))
	fmt.Printf("Server Running")
	http.ListenAndServe(":4352", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "hello.html", nil)
}

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do your cors stuff
		// return if you do not want the FileServer handle a specific request
		enableCors(&w)
		fs.ServeHTTP(w, r)
	}
}
