package main

import (
    "fmt"
    "html/template"
    "net/http"
)

// Model
type Model struct {
    Message string
}

// View
func renderView(w http.ResponseWriter, m Model) {
    tmpl, err := template.New("index").Parse(`<html><body><h1>{{.Message}}</h1></body></html>`)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, m)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// Presenter
func helloHandler(w http.ResponseWriter, r *http.Request) {
    model := Model{Message: "Hello, World!"}
    renderView(w, model)
}

func main() {
    http.HandleFunc("/", helloHandler)

    port := 8080
    fmt.Printf("Listening on :%d...\n", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}