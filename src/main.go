package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	logger  *Logger
	storage *Storage
)

const (
	webAddress = "localhost:8083"
)

func main() {
	logger = NewLogger(os.Stdout, os.Stdout, os.Stdout, llInfo)
	storage = NewStorage(OpenDBConfig("./dbconfig.json"))

	defer storage.Close()

	storage.CreateTables()

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	router.PathPrefix("/forms/").Handler(http.StripPrefix("/forms/",
		http.FileServer(http.Dir("./forms/"))))

	router.HandleFunc("/", HandlerBase)
	router.HandleFunc("/base", HandlerBase)

	router.HandleFunc("/clients", HandlerClients)
	router.HandleFunc("/calculate", HandlerCalculate)
	router.HandleFunc("/compony", HandlerCompony)
	router.HandleFunc("/calucaltions", HandlerCalculation)

	router.HandleFunc("/add_client", HandlerAddClient)
	router.HandleFunc("/add_counter_info", HandlerAddCounterInfo)
	router.HandleFunc("/add_counter", HandlerAddCounter)
	router.HandleFunc("/add_flat", HandlerAddFlat)
	router.HandleFunc("/add_compony", HandlerAddCompony)

	router.HandleFunc("/show_componies", HandlerShowComponies)

	logger.Info(webAddress)

	if err := http.ListenAndServe(webAddress, router); err != nil {
		logger.Error(err.Error())
		panic(err)
	}
}
