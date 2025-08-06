# Internal Transfers System

This Go-based application facilitates internal financial transfers between accounts using a RESTful HTTP interface and PostgreSQL for data persistence.

---

## 🔧 Requirements

- Go 1.20+
- PostgreSQL (running locally or in Docker)
  Database name: transfers

---

##  Setup Instructions

1. **Clone the repo:**
   
   git clone https://github.com/Rashidali23/Internal-transfers-System.git
   cd Internal-transfers-System

 2. **Run the application:**
    go run main.go
  NOTE: Make sure PostgreSQL is running and accessible

##  Testing Instructions

Use Postman or curl to test:

 Create Account:
                ```curl -X POST http://localhost:8080/accounts -H "Content-Type: application/json" \
-d '{"account_id": 1, "initial_balance": "100.00000"}'
```



Get Account:
        ``` curl http://localhost:8080/accounts/1 
```


Make Transaction:
            ``` curl -X POST http://localhost:8080/transactions -H "Content-Type: application/json" \
-d '{"source_account_id": 1, "destination_account_id": 2, "amount": "20.00000"}'  
```
