
## Task
-------------------------

Terms of Reference for GO Developer
1. Write two micro-services:
- Name: Accounts
- Name: Bank
2. Micro-service Accounts, must:
- Have a server architecture and implement the following in protobuf v3


Account structure with fields: ID (generate a unique UUID upon creation), walletID of string type, balance of type uint64

### Services:
1. Create Account (creates a new account, walletID must be empty, balance 0)
2. Get Accounts (returns accounts) stream
3. Generate Address (generates random UUID and writes to walletID)
4. Deposit (adds the amount to balance, which is passed to request) returns account
5. Withdrawal (subtracts the amount from balance, which is passed to request) returns account


Repository/Store must be implemented in Mongo Logs on Zap
3. Micro-service Bank:
- Uses the same protobuf
- Each service from Accounts converts to REST endpoints written in Gin. Those. is a proxy service for Accounts.
- Calls the appropriate rpc services on the Accounts micro service.

## Installation
---
```bash
cd bank-and-accounts
go mod tidy
```

## Run Account and Bank Services
----
```bash
go run accounts/main.go
go run bank/main.go
```

## Usage
---
```bash
# get list of accounts
curl -L -X GET 'localhost:8085/account'

# create new account
curl -L -X POST 'localhost:8085/account' -H 'Content-Type: application/json' -d '{}'

# (re)generate walletID to the certain account
curl -L -X POST 'localhost:8085/account/generate-address' -H 'Content-Type: application/json' -d '{"id":"99dcc6ff-3a07-4c31-a5a5-901bb67ea933"}'

# deposit to the certain account
curl -L -X POST 'localhost:8085/account/deposit' -H 'Content-Type: application/json' -d '{"id":"99dcc6ff-3a07-4c31-a5a5-901bb67ea933","amount":5}'

# withdrawal from the certain account
curl -L -X POST 'localhost:8085/account/withdrawal' -H 'Content-Type: application/json' -d '{"id":"99dcc6ff-3a07-4c31-a5a5-901bb67ea933","amount":2}'
```