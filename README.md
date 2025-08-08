# Internal Transfers System

This Go-based application facilitates internal financial transfers between accounts using a RESTful HTTP interface and PostgreSQL for data persistence.

---

## 🔧 Requirements

- Go 1.20+
- PostgreSQL (running locally or in Docker)

## 🔧 Requirements
 Before running the application, you need to configure your PostgreSQL connection details.

 Update your conf.env file with correct PostgreSQL credentials:
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=user
    DB_NAME=transfers
    DB_MASTER_NAME=postgres
    DB_SSLMODE=disable

DB_HOST: Hostname or IP where your PostgreSQL is running

DB_PORT: Port number (default is 5432)

DB_USER: PostgreSQL username

DB_PASSWORD: PostgreSQL password

DB_NAME: The database your app will use (e.g., transfers) 

DB_MASTER_NAME: The default database used to create DB_NAME if it doesn't exist (usually postgres)

DB_SSLMODE: Set to disable for local setups

⚠️ Make sure these values match your PostgreSQL setup, otherwise the app won't connect correctly.
---

##  Setup Instructions

1. **Clone the repo:**
   
   git clone https://github.com/Rashidali23/Internal-transfers-System.git
   cd Internal-transfers-System

 2. **Run the application:**
   NOTE:Ensure PostgreSQL is running and the conf.env file is correctly configured.
    go run main.go

##  Testing Instructions

Use Postman or curl to test:

 Create Account:
               curl -X POST http://localhost:8080/accounts -H "Content-Type: application/json" \-d '{"account_id": 1, "initial_balance": "100.00000"}'




Get Account:
         curl http://localhost:8080/accounts/1 



Make Transaction:
             curl -X POST http://localhost:8080/transactions -H "Content-Type: application/json" \-d '{"source_account_id": 1, "destination_account_id": 2, "amount": "20.00000"}'  

