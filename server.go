package main

import (
  "net/http"
  "html/template"
)

// for use with templates
type Page struct {
  Name string
  Title string
}

const viewsPath = "/views/"

var templates = template.Must(template.ParseGlob("templates/*.html"))

func handle(t string) (string, http.HandlerFunc) {
  return viewsPath+t, func(w http.ResponseWriter, r *http.Request) {
    p := &Page{ Name: "Reed", Title: t }
    err := templates.ExecuteTemplate(w, t+".html", p)
    if err != nil {
      http.NotFound(w, r)
    }
  }
}

func main() {
  //TODO leave main blank and make it a FIXME
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
  http.HandleFunc(handle("gophers"))
  http.HandleFunc(handle("home"))
  http.ListenAndServe(":8080", nil)
}
