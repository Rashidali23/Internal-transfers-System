package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=postgres password=user dbname=transfers sslmode=disable"
	var err error
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
