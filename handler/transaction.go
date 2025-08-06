package handler

import (
	"Internal-transfers-System/db"
	"Internal-transfers-System/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tx models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(tx.Amount,64)
	if err != nil{
		log.Panicln("Error while converting to float")
	}
	if amount <= 0 {
		http.Error(w, "Amount must be positive", http.StatusBadRequest)
		return
	}

	txFunc := func(txn *sql.Tx) error {
		var sourceBal float64
		err := txn.QueryRow("SELECT balance FROM accounts WHERE account_id = $1 FOR UPDATE", tx.SourceAccountID).Scan(&sourceBal)
		if err != nil {
			return fmt.Errorf("source account error: %v", err)
		}

		if sourceBal < amount {
			return fmt.Errorf("insufficient funds")
		}

		_, err = txn.Exec("UPDATE accounts SET balance = balance - $1 WHERE account_id = $2", tx.Amount, tx.SourceAccountID)
		if err != nil {
			return err
		}

		_, err = txn.Exec("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2", tx.Amount, tx.DestinationAccountID)
		if err != nil {
			return err
		}

		_, err = txn.Exec("INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)", tx.SourceAccountID, tx.DestinationAccountID, tx.Amount)
		return err
	}

	txn, err := db.DB.Begin()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	err = txFunc(txn)
	if err != nil {
		txn.Rollback()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	txn.Commit()
	w.WriteHeader(http.StatusCreated)
}