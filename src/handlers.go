package main

import (
	"html/template"
	"net/http"
)

func HandlerBase(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerAddClient(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/html/clients.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}
