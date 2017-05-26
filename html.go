package main

import (
	"html/template"
	"net/http"
)

var varmap map[string]interface{}

// ServiHTML fa partire il server html
func ServiHTML() {
	varmap = map[string]interface{}{
		"matrice":       Matrix,
		"tempoAggiorna": Clock,
		"larghezza":     Larghezza,
		"altezza":       Altezza,
		
	}
	http.HandleFunc("/tabella", handlerRoot("template/tabella.html"))
	http.HandleFunc("/", handlerRoot("template/Interfaccia.html"))
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":3000", nil)
}

func handlerRoot(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles(path)
		if err != nil {
			panic(err.Error())
		}
		templ.Execute(w, varmap)
	}
}
