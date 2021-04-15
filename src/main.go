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

	router.HandleFunc("/", HandlerBase)
	router.HandleFunc("/base", HandlerBase)

	router.HandleFunc("/add_client", HandlerAddClient)

	logger.Info(webAddress)

	if err := http.ListenAndServe(webAddress, router); err != nil {
		logger.Error(err.Error())
		panic(err)
	}
}
