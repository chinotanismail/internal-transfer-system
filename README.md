# Internal Transfer System

This project is a wallet microservice designed to manage account balances and perform balance transactions. It is intended to be used as part of a larger system, where detailed account information is managed by a separate service. The microservice provides APIs for balance management and transaction processing, and maintains a ledger of all transactions.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher recommended)
- [PostgreSQL](https://www.postgresql.org/download/) (version 12 or higher recommended)

## Setup Instructions

### 1. Clone the Repository

```bash
git clone <repository-url>
cd internal-transfer-system
```

### 2. Create the PostgreSQL Database

Create a PostgreSQL database named `internal_transfer_system` (or another name of your choice):

```bash
psql -U <your_postgres_user> -c "CREATE DATABASE internal_transfer_system;"
```

Replace `<your_postgres_user>` with your PostgreSQL username.

### 3. Configure the Database Connection

Edit the DSN (Data Source Name) in `config/db.go` to match your PostgreSQL setup. The DSN typically looks like:

```go
dsn := "host=localhost user=postgres password=yourpassword dbname=internal_transfer_system port=5432 sslmode=disable"
```

- Change `user`, `password`, `dbname`, and other parameters as needed for your environment.

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run the Service

The service will automatically migrate (create) the required tables in the database upon startup.

```bash
go run main.go
```

The service will start and listen on the configured port (see `main.go` for details).

## Project Assumptions

- **Wallet Microservice Scope:**  
  This service is designed to handle only account balances and perform balance transactions. All other detailed information about accounts (such as user profiles, authentication, etc.) is managed by another service.

- **Transactions Table as Ledger and History:**  
  The `transactions` table serves as a ledger for all balance changes. It can also be used to track and debug the history of attempted and completed transactions for accounts.

- **Transactions IDs are sequential:**  
  The `transactions` table primary key (ID) also serves as its transaction ID and is sequential instead of generated.
## API Endpoints

### `POST /accounts`
Create a new account.

- **Request Body:** JSON object with required account details:
  ```json
  {
    "account_id": 123,
    "initial_balance": "100.00"
  }
  ```
- **Success Response:** HTTP 200 OK with an empty body.
- **Error Response:** HTTP 400/500 with JSON:
  ```json
  { "error": "Error message" }
  ```

### `GET /accounts/:account_id`
Retrieve account information by account ID.

- **Path Parameter:** `account_id` (string or integer, depending on implementation)
- **Response:** JSON object with the account's details and current balance.

### `POST /transactions`
Create a new transaction between accounts.

- **Request Body:** JSON object specifying source account, destination account, and amount:
  ```json
  {
    "source_account_id": 1,
    "destination_account_id": 2,
    "amount": "50.00"
  }
  ```
- **Success Response:** HTTP 200 OK with JSON:
  ```json
  {
    "message": "Transaction created successfully",
    "transaction": {
      "id": 1,
      "source_account_id": 1,
      "destination_Account_id": 2,
      "amount": 50.0,
      "created_at":"2025-06-15T19:33:17.99096+07:00",
      "status": "success"
    }
  }
  ```
- **Error Response:** HTTP 400/500 with JSON:
  ```json
  { "error": "Error message" }
  ```
