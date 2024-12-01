# Voucher Management API Service

## Project Overview
This is a Golang-based API service for managing vouchers in an e-commerce application. The project allows for the creation, management, and redemption of two types of vouchers: e-commerce vouchers and point redemption vouchers.
### Team Members
- Muhammad Siskandar Zulkarnain
- Dewi Safira Ramadhani

### Task Manager
https://docs.google.com/spreadsheets/d/12VOdT_kfzBBAEX_SW9Wyzh-3Z6Rkta5nyZa4Sx-VNZM/edit?usp=sharing

## Project Features

### Voucher Management
- Create new vouchers with detailed specifications
- Delete vouchers
- Edit voucher details
- View voucher lists with advanced filtering

### Voucher Types
1. E-commerce Vouchers
   - Free Shipping
   - Percentage or Nominal Discounts
   - Specific usage conditions

2. Point Redemption Vouchers
   - Redeem vouchers using user points
   - Track redemption history

## Technical Stack
- Language: Golang
- Framework: Gin (REST API)
- Database: PostgreSQL
- ORM: GORM
- Logging: Zap Logger

## Prerequisites
- Golang 1.16+
- PostgreSQL
- Postman (for API testing)

## Installation

### Clone the Repository
```
git clone https://github.com/yourusername/project-sistem-voucher-golang-homework-Muhamad-Siskandar-Zulkarnain-Dewi-Safira-Ramadhani.git
cd project-sistem-voucher-golang-homework-Muhamad-Siskandar-Zulkarnain-Dewi-Safira-Ramadhani
```

## Setup Environment
1. Install dependencies
```
go mod tidy
```
2. Configure Database
   - Create a PostgreSQL database
   - Update database configuration in .env or configuration file
3. Run Migrations
  ```
  go run migrate.go 
  ```
4. Run Seeder (if applicable)
   ```
   go run seeder.go
   ```
## Running the Application
```
go run main.go
```

## API Endpoints
### Voucher Management
- POST /vouchers - Create new voucher
- DELETE /vouchers/:id - Delete voucher
- PUT /vouchers/:id - Update voucher details
- GET /vouchers - List vouchers with filters
### Voucher Redemption
- GET /vouchers/redeemable - List redeemable vouchers
- POST /vouchers/redeem - Redeem a voucher
- GET /vouchers/user - User's voucher list
### Voucher Usage
- POST /vouchers/validate - Validate voucher
- POST /vouchers/use - Use a voucher
  
### Testing
- Run unit tests
```
go test ./... -cover
```

## Team Members
- Muhammad Siskandar Zulkarnain
- Dewi Safira Ramadhani

## Task Manager
https://docs.google.com/spreadsheets/d/12VOdT_kfzBBAEX_SW9Wyzh-3Z6Rkta5nyZa4Sx-VNZM/edit?usp=sharing
