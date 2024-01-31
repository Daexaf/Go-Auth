# PaymentAPI

PaymentAPI is a Go (Golang) project for handling payment transactions.

## Features
- User authentication (signup, login, logout)
- Transaction handling (create transactions)
- User validation

## Getting Started

### Prerequisites
- Go (Golang) installed
- PostgreSQL installed and running
- Git installed

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/PaymentAPI.git
   cd PaymentAPI
2. go mod tidy
3. go run main.go

### Authentication
POST /login
{
  "email": "user@example.com",
  "password": "yourpassword"
}

# Response
{
  "token": "your-jwt-token"
}

### Create Transaction
POST /create-transaction
Content-Type: application/json
Authorization: Bearer your-jwt-token

{
  "SenderAccountID": 1,
  "ReceiveAccountID": 2,
  "Amount": "500.00"
}

# Response
{
  "message": "Transaction completed"
}


