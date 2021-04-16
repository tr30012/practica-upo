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

	r.ParseForm()
	logger.Info(r.Form)

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

	date := r.FormValue("date")
	fio := r.FormValue("fio")
	pass := r.FormValue("pass")
	sex := r.FormValue("sex")
	tel := r.FormValue("sex")

	storage.ExecuteString(SQLInsertClient, date, tel, fio, pass, sex)

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerAddCounterInfo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/clients", http.StatusSeeOther)
}

func HandlerAddCounter(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./forms/add_counter_info.html")

	r.ParseForm()
	logger.Info(r.Form)

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerAddFlat(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./forms/add_counter.html")

	/*
		city := r.FormValue("city")
		compony := r.FormValue("compony")
		flat := r.FormValue("flat")
		house := r.FormValue("hause")
		sq := r.FormValue("sq")
		street := r.FormValue("street")
		t := r.FormValue("t")
	*/

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func HandlerAddCompony(w http.ResponseWriter, r *http.Request) {
	cost := r.FormValue("cost")
	fio := r.FormValue("fio")
	name := r.FormValue("name")
	recs := r.FormValue("recs")

	storage.ExecuteString(SQLInsertCompony, name, cost, recs, fio)

	http.Redirect(w, r, "/compony", http.StatusSeeOther)
}

func HandlerShowComponies(w http.ResponseWriter, r *http.Request) {
	rows := storage.QueryString(SQLSelectAllCompanies)

	companies := make([]Company, 0)

	for rows.Next() {
		var compony Company
		rows.Scan(&compony.Id, &compony.Name, &compony.Cost, &compony.Recs, &compony.FIO)
		companies = append(companies, compony)
	}

	if err := rows.Close(); err != nil {
		logger.Error(err)
		panic(err)
	}

	t, err := template.ParseFiles("./templates/show_componies.html")

	if err != nil {
		http.Error(w, err.Error(), 1)
		logger.Error(err.Error())
	}

	t.Execute(w, companies)
}
