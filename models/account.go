package models

type Account struct {
	AccountID int    `json:"account_id"`
	Balance   string `json:"initial_balance"`
}