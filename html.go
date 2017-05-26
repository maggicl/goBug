package main

import (
	"html/template"
	"net/http"
)

// ServiHTML fa partire il server html
func ServiHTML() {
	http.HandleFunc("/", handlerRoot)
	http.ListenAndServe(":3000", nil)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("template/Interfaccia.html")
	if err != nil {
		panic(err.Error())
	}
	templ.Execute(w, new(interface{}))
}
