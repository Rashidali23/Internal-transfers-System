package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"Internal-transfers-System/db"

	"github.com/gorilla/mux"
	"Internal-transfers-System/models"
)


func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO accounts (account_id, balance) VALUES ($1, $2)", acc.AccountID, acc.Balance)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating account: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var acc models.Account
	err = db.DB.QueryRow("SELECT account_id, balance FROM accounts WHERE account_id = $1", id).Scan(&acc.AccountID, &acc.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Account not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error fetching account: %v", err), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(acc)
}