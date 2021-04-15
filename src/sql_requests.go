package main

const (
	SQLCreateCamponyTable = `
		CREATE TABLE IF NOT EXISTS Company(
			ID_company INT NOT NULL PRIMARY KEY,
			Nazv VARCHAR(20) NOT NULL,
			Tariff INT NOT NULL,
			Recvizit VARCHAR(30) NOT  NULL,
			FIO_Ruck VARCHAR(40) NOT  NULL
			);		
	`

	SQLCreateSobstvTable = `
		CREATE TABLE IF NOT EXISTS Sobstv(
			ID_sobstv INT NOT NULL PRIMARY KEY,
			Data_birthday DATE NOT NULL,
			Phone VARCHAR(12) NOT  NULL,
			FIO_sobstv VARCHAR(40) NOT  NULL,
			Passport VARCHAR(10) NOT NULL,
			Gender VARCHAR(1) NOT NULL
			);		
	`

	SQLCreateFlatTable = `
		CREATE TABLE IF NOT EXISTS Flat(
			ID_flat INT NOT NULL PRIMARY KEY,
			Adress VARCHAR(40) NOT NULL,
			Other_info VARCHAR(40) NOT  NULL,
			Temperatura VARCHAR(3) NOT NULL,
			ID_sobstv INT NOT NULL REFERENCES Sobstv(ID_sobstv),
			ID_company INT NOT NULL REFERENCES Company(ID_company)
			);		
	`

	SQLCreateLgotTable = `
		CREATE TABLE IF NOT EXISTS Lgot(
			ID_lgot INT NOT NULL PRIMARY KEY,
			ID_sobstv INT NOT NULL REFERENCES Sobstv(ID_sobstv),
			Discount VARCHAR(4) NOT  NULL
			);	
	`

	SQLCreateSchetchikTable = `
		CREATE TABLE IF NOT EXISTS Schetchik(
			ID_schetchik INT NOT NULL PRIMARY KEY,
			ID_flat INT NOT NULL REFERENCES Flat(ID_flat),
			Date_ustanovk DATE NOT  NULL,
			Firma VARCHAR(25) NOT NULL
			);		
	`

	SQLCreatePokazTable = `
		CREATE TABLE IF NOT EXISTS Pokaz(
			ID_pokaz INT NOT NULL PRIMARY KEY,
			Adout_pokaz VARCHAR(10) NOT NULL,
			ID_schetchik INT NOT NULL REFERENCES Schetchik(ID_schetchik),
			Month_pokaz VARCHAR(8) NOT NULL
			);
	`

	SQLCreatePayTable = `
		CREATE TABLE IF NOT EXISTS Pay(
			ID_pay INT NOT NULL PRIMARY KEY,
			Data_pay DATE NOT NULL,
			All_summa VARCHAR(10) NOT  NULL,
			ID_company INT NOT NULL REFERENCES Company(ID_company),
			ID_lgot INT NOT NULL REFERENCES Lgot(ID_lgot),
			ID_sobstv INT NOT NULL REFERENCES Sobstv(ID_sobstv),
			ID_flat INT NOT NULL REFERENCES Flat(ID_flat),
			ID_schetchik INT NOT NULL REFERENCES Schetchik(ID_schetchik)
			);
	`
)
