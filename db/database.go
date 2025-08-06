package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "host=localhost port=5432 user=postgres password=user dbname=transfers sslmode=disable"
const masterConnStr = "host=localhost port=5432 user=postgres password=user dbname=postgres sslmode=disable"

var DB *sql.DB

func InitDB() {
	var err error

	// Step 1: Connect to the default "postgres" database
	masterDB, err := sql.Open("postgres", masterConnStr)
	if err != nil {
		log.Fatalf("Error connecting to master DB: %v", err)
	}
	defer masterDB.Close()

	// Step 2: Check if the "transfers" database exists, create if not
	_, err = masterDB.Exec("CREATE DATABASE transfers")
	if err != nil && err.Error() != `pq: database "transfers" already exists` {
		log.Fatalf("Error creating transfers DB: %v", err)
	}
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	createTables()
}

func createTables() {
	accountTable := `CREATE TABLE IF NOT EXISTS accounts (
		account_id INT PRIMARY KEY,
		balance NUMERIC
	);`

	transactionTable := `CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		source_account_id INT,
		destination_account_id INT,
		amount NUMERIC,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(accountTable)
	if err != nil {
		log.Fatalf("Error creating accounts table: %v", err)
	}
	_, err = DB.Exec(transactionTable)
	if err != nil {
		log.Fatalf("Error creating transactions table: %v", err)
	}
}
