package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"html/template"
)

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/env", env)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func welcome(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("tmpl/welcome.html"))
	t.Execute(w, "Welcome to Golang Predix Microservice Template")
}

func env(w http.ResponseWriter, r *http.Request) {
	out := "<dl>"
	for _, env := range os.Environ() {
		kv := strings.Split(env, "=")
		out += "<dt>" + kv[0] + "</dt>"
		out += "<dd>" + kv[1] + "</dd>"
	}
	out += "</dl>"
	fmt.Fprintln(w, `<div class="envs">`+out+`</div>`)
}
