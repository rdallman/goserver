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
  p := &Page{}  //FIXME to include your name
  t, _ := template.ParseFiles("index.html")
  t.Execute(w, p)
}

func main() {
  //TODO leave main blank and make it a FIXME
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":8080", nil)
}
