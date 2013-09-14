package main

import (
  "net/http"
  "html/template"
)

// for use with templates
type Page struct {
  Name string
}

// for '/'
func indexHandler(w http.ResponseWriter, r *http.Request) {
  p := &Page{Name: "Reed"}  //FIXME to include your name
  t, _ := template.ParseFiles("index.html")
  t.Execute(w, p)
}

// for '/views/gopher'
func gopherHandler(w http.ResponseWriter, r *http.Request) {
  p := &Page{ }
  t, _ := template.ParseFiles("gophers.html")
  t.Execute(w, p)
}

func main() {
  //TODO leave main blank and make it a FIXME
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/views/gopher", indexHandler)
  http.ListenAndServe(":8080", nil)
}
