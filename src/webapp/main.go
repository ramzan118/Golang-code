package main

import (
    "html/template"
    "log"
    "net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles("templates/" + tmpl + ".html")
    if err != nil {
        log.Fatal(err)
    }
    t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "form")
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    // Handle form submission logic here
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/submit", submitHandler)
    log.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
