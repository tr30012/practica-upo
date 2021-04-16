package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	Connection string `json:"connection"`
	Driver     string `json:"driver"`
}

func OpenDBConfig(path string) *DBConfig {
	logger.Info("Окрытие конфига из файла")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Warning("Файл не найден. Используется стандартный конфиг")

		return OpenDBConfigDefault()
	}

	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	fContent, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	var cfg DBConfig
	json.Unmarshal(fContent, &cfg)

	defer f.Close()
	return &cfg
}

func OpenDBConfigDefault() *DBConfig {
	return &DBConfig{
		Connection: "./storage/practica-upo.db3",
		Driver:     "sqlite3",
	}
}

type Storage struct {
	db *sql.DB
}

func NewStorage(cfg *DBConfig) *Storage {
	logger.Info("Попытка создать соединение <", cfg.Connection, ">")
	connection, err := sql.Open(cfg.Driver, cfg.Connection)

	if err != nil {
		logger.Error("Ошибка соединения: ", err.Error())
		panic(err)
	}

	logger.Info("Проверка соединения <", cfg.Connection, ">")
	if err := connection.Ping(); err != nil {
		logger.Error("Ошибка соединения: ", err.Error())
		panic(err)
	}

	storage := Storage{
		db: connection,
	}

	logger.Info("Успешно")

	return &storage
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) CreateTables() {
	s.ExecuteString(SQLCreateCamponyTable)
	s.ExecuteString(SQLCreateSobstvTable)
	s.ExecuteString(SQLCreateFlatTable)
	s.ExecuteString(SQLCreateLgotTable)
	s.ExecuteString(SQLCreateSchetchikTable)
	s.ExecuteString(SQLCreatePokazTable)
	s.ExecuteString(SQLCreatePayTable)
}

func (s *Storage) ExecuteString(cmd string, a ...interface{}) {
	q, err := s.db.Prepare(cmd)
	logger.Info("Подготовка строки: ", cmd)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	if _, err := q.Exec(a...); err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	logger.Info("Успешно")
}

func (s *Storage) QueryString(cmd string, a ...interface{}) *sql.Rows {
	logger.Info("Подготовка строки: ", cmd)
	rows, err := s.db.Query(cmd, a...)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	logger.Info("Успешно")
	return rows
}

func (s *Storage) GetComponyByName(name string) (Company, error) {
	rows := storage.QueryString(SQLSelectAllCompanies)

	var compony Company

	for rows.Next() {

		rows.Scan(&compony.Id, &compony.Name, &compony.Cost, &compony.Recs, &compony.FIO)

		if name == compony.Name {

			if err := rows.Close(); err != nil {
				logger.Error(err.Error())
				panic(err)
			}

			return compony, nil
		}
	}

	if err := rows.Close(); err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	return compony, errors.New("nil")
}
