# 📈 Go Retail ISA Investment Service

A backend service that allows retail customers to invest in ISAs by selecting funds and specifying investment amounts.

## 🌟 Overview
This service enables X platform's retail (direct) customers to:

- 📜 View available investment funds
- ✅ Select a single fund
- 💰 Specify an investment amount
- 🔍 Record these selections for later querying

The solution is built using Go, following a clean architecture approach with distinct layers:

- **Domain Layer**: Core business entities and interfaces
- **Service Layer**: Business logic implementation
- **Repository Layer**: Data storage and retrieval
- **API Layer**: HTTP interface for clients

## 🏗 Architecture
### Clean Architecture Principles
This project follows clean architecture principles to ensure:

- 🛠 Separation of concerns
- 🧪 Testability
- 📈 Extensibility
- 📌 Domain-driven design

### 📂 Project Structure
```
cushon-retail-isa-service/
├── cmd/
│   └── api/            # Application entry point
│       └── main.go
├── internal/
│   ├── domain/         # Core business entities and interfaces
│   ├── service/        # Business logic implementation
│   ├── repository/     # Data storage implementations 
│   └── api/
│       └── handler/    # HTTP handlers for API endpoints
```

## 📌 Domain Models
This service defines three primary domain entities:

- **👤 Customer**: Represents a retail customer investing in ISAs
- **📊 Fund**: Represents an investment fund option
- **💵 Investment**: Represents a customer's investment selection in a specific fund

## 🔑 Key Design Decisions
### 1️⃣ Service Layer Pattern
- Encapsulates all business rules
- Depends on repository interfaces for data access
- Provides a clean separation between business rules and data access

### 2️⃣ Money Handling
- Monetary values stored as integers (pence) to avoid floating-point precision issues
- Conversion between display format (£) and storage format (pence) happens at the API boundary

```go
// Convert input from API (e.g., "25000.00")
amountFloat, err := strconv.ParseFloat(req.Amount, 64)  // Convert to float
amountPence := int64(amountFloat * 100)                 // Convert to pence

// Convert back when returning to client
response.AmountValue = float64(investment.Amount) / 100.0
```

### 3️⃣ Repository Pattern
- Repository interfaces in the domain layer define data access methods
- Current implementation uses in-memory storage for simplicity
- Interfaces allow for easy replacement with database implementations

### 4️⃣ Error Handling
- 🛑 Domain-specific error types
- ✅ Validation before state changes
- 📡 Appropriate HTTP status codes at the API layer

### 5️⃣ Graceful Shutdown
- Captures termination signals (CTRL+C, kill commands)
- Completes in-flight requests before shutting down
- Uses a timeout to prevent hanging indefinitely

## 🔥 API Usage
### 🚀 Getting Started
Run the application:
```bash
go run cmd/api/main.go
```
The server will start on port `8080` by default with seeded test data.

### 🔗 Example API Requests
#### 📌 List All Available Funds
```bash
curl -X GET http://localhost:8080/api/v1/funds | jq
```
#### 📌 Create an Investment
```bash
curl -X POST http://localhost:8080/api/v1/investments \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "customer-1",
    "fund_id": "fund-1",
    "amount": "15000.00"
  }' | jq
```
#### 📌 Get an Investment
```bash
curl -X GET http://localhost:8080/api/v1/investments/inv-123abc | jq
```

#### ⚠️ Error Example: Exceeding ISA Limit (£20,000)
```bash
curl -X POST http://localhost:8080/api/v1/investments \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "customer-1",
    "fund_id": "fund-1",
    "amount": "25000.00"
  }'
```
**Response:**
```
investment exceeds ISA annual limit of £20,000
```

## 🔍 Assumptions
- **🛡 Authentication & Authorization**: To be handled by middleware/gateway
- **📅 ISA Regulatory Compliance**: Basic ISA limit validation (£20,000 annually)
- **📊 Single Fund Selection**: Customers can select only one fund per investment
- **👥 Customer Onboarding**: Assumes customers already exist in the system

## 🚀 Future Improvements
### 📦 Data Persistence
- Implement **PostgreSQL** for storage
- Add database migrations for schema evolution

### 📊 Multiple Fund Selection
- Add **investment_items** table for allocations
- Implement percentage-based allocation logic

### 🔐 Security Enhancements
- JWT-based authentication
- Role-based access control
- Audit logging for financial transactions

### 📈 Operational Readiness
- Docker support
- Observability with **metrics & structured logging**
- Add **health check endpoints**

## 🛠 Testing Strategy
The current implementation includes basic **unit tests** focusing on business logic and ISA limit validation.

### ✅ Future Testing Improvements
- **Unit Tests**: Service layer coverage, repository tests, edge cases
- **Integration Tests**: API endpoint tests, mock external service integrations
- **End-to-End Tests**: Complete user journeys, performance testing
- **CI/CD Integration**: Automated test runs, test coverage enforcement

---
