package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load("conf.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read env variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	masterDBName := os.Getenv("DB_MASTER_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// Prepare connection strings
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	masterConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, masterDBName, sslmode)

	// Step 1: Connect to the default "postgres" database
	masterDB, err := sql.Open("postgres", masterConnStr)
	if err != nil {
		log.Fatalf("Error connecting to master DB: %v", err)
	}
	defer masterDB.Close()

	// Step 2: Create the "transfers" database if not exists
	_, err = masterDB.Exec("CREATE DATABASE " + dbname)
	if err != nil && err.Error() != fmt.Sprintf(`pq: database "%s" already exists`, dbname) {
		log.Fatalf("Error creating transfers DB: %v", err)
	}

	// Step 3: Connect to the "transfers" DB
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
