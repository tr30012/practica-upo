package main

const (
	SQLCreateCamponyTable = `
		CREATE TABLE IF NOT EXISTS Company(
			ID_company INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Nazv VARCHAR(20) NOT NULL,
			Tariff INTEGER NOT NULL,
			Recvizit VARCHAR(30) NOT  NULL,
			FIO_Ruck VARCHAR(40) NOT  NULL
			);		
	`

	SQLCreateSobstvTable = `
		CREATE TABLE IF NOT EXISTS Sobstv(
			ID_sobstv INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Data_birthday DATE NOT NULL,
			Phone VARCHAR(12) NOT  NULL,
			FIO_sobstv VARCHAR(40) NOT  NULL,
			Passport VARCHAR(10) NOT NULL,
			Gender VARCHAR(1) NOT NULL
			);		
	`

	SQLCreateFlatTable = `
		CREATE TABLE IF NOT EXISTS Flat(
			ID_flat INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Adress VARCHAR(40) NOT NULL,
			Other_info VARCHAR(40) NOT  NULL,
			Temperatura VARCHAR(3) NOT NULL,
			ID_sobstv INTEGER NOT NULL REFERENCES Sobstv(ID_sobstv),
			ID_company INTEGER NOT NULL REFERENCES Company(ID_company)
			);		
	`

	SQLCreateLgotTable = `
		CREATE TABLE IF NOT EXISTS Lgot(
			ID_lgot INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			ID_sobstv INTEGER NOT NULL REFERENCES Sobstv(ID_sobstv),
			Discount VARCHAR(4) NOT  NULL
			);	
	`

	SQLCreateSchetchikTable = `
		CREATE TABLE IF NOT EXISTS Schetchik(
			ID_schetchik INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			ID_flat INTEGER NOT NULL REFERENCES Flat(ID_flat),
			Date_ustanovk DATE NOT  NULL,
			Firma VARCHAR(25) NOT NULL
			);		
	`

	SQLCreatePokazTable = `
		CREATE TABLE IF NOT EXISTS Pokaz(
			ID_pokaz INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Adout_pokaz VARCHAR(10) NOT NULL,
			ID_schetchik INTEGER NOT NULL REFERENCES Schetchik(ID_schetchik),
			Month_pokaz VARCHAR(8) NOT NULL
			);
	`

	SQLCreatePayTable = `
		CREATE TABLE IF NOT EXISTS Pay(
			ID_pay INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Data_pay DATE NOT NULL,
			All_summa VARCHAR(10) NOT  NULL,
			ID_company INTEGER NOT NULL REFERENCES Company(ID_company),
			ID_lgot INTEGER NOT NULL REFERENCES Lgot(ID_lgot),
			ID_sobstv INTEGER NOT NULL REFERENCES Sobstv(ID_sobstv),
			ID_flat INTEGER NOT NULL REFERENCES Flat(ID_flat),
			ID_schetchik INTEGER NOT NULL REFERENCES Schetchik(ID_schetchik)
			);
	`

	SQLInsertClient  = `INSERT INTO Sobstv(Data_birthday, Phone, FIO_sobstv, Passport, Gender) VALUES (?, ?, ?, ?, ?)`
	SQLInsertFlat    = `INSERT INTO Flat(Adress, Other_info, Temperatura, ID_sobstv, ID_company) VALUES (?, ?, ?, ?, ?)`
	SQLInsertCompony = `INSERT INTO Company(Nazv, Tariff, Recvizit, FIO_Ruck) VALUES (?, ?, ?, ?)`

	SQLSelectAllCompanies = `SELECT * FROM Company`
)
