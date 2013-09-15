package main

import (
  "net/http"
  "html/template"
)

// for use with templates
type Page struct {
  Name string
}

// find me at "/"
func indexHandler(w http.ResponseWriter, r *http.Request) {
  p := &Page{ Name: "" }  //FIXME to include your name
  t, _ := template.ParseFiles("index.html")
  t.Execute(w, p)
}

func main() {
  //FIXME
}
