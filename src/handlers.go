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

func HandlerClients(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/clients.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerCalculate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/calculate.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerCompony(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/compony.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerCalculation(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/calculations.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerAddClient(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./forms/add_flat.html")

	r.ParseForm()
	logger.Info(r.Form)

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerAddCounterInfo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/clients.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerAddCounter(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./forms/add_counter_info.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerAddFlat(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./forms/add_counter.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}
