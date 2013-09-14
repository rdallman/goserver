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

// for '/'
func indexHandler(w http.ResponseWriter, r *http.Request) {
  p := &Page{Name: "Reed"}  //FIXME to include your name
  t, _ := template.ParseFiles("templates/index.html") //FIXME back to index.html
  t.Execute(w, p)
}

// for '/views/gopher'
func gopherHandler(w http.ResponseWriter, r *http.Request) {
  renderTemplate(w, r, "gophers")
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplate(w http.ResponseWriter, r *http.Request, t string) {
  p := &Page{ Name: "Reed", Title: t }
  err := templates.ExecuteTemplate(w, t+".html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func main() {
  //TODO leave main blank and make it a FIXME
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/views/gophers", gopherHandler)
  http.ListenAndServe(":8080", nil)
}
