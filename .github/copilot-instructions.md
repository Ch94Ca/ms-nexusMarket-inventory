# Copilot Instructions for ms-nexusMarket-inventory

## Project Overview
This is a Go-based inventory microservice implementing hexagonal (clean) architecture with multiple persistence layers and asynchronous processing. The service handles stock reservations, releases, and queries with complete audit trails.

## Architecture Patterns

### Hexagonal Architecture
- `internal/domain/`: Pure business entities and interfaces (Stock, Product models)
- `internal/app/`: Application layer with use cases and orchestration logic
- `internal/infra/`: Infrastructure adapters (PostgreSQL, MongoDB, Kafka clients)
- `internal/strategy/`: Strategy pattern implementations for different reservation policies
- `cmd/api/`: Application entry point (main.go for HTTP server)

### Key Components
- **Primary ports**: REST API endpoints (`/stock/reserve`, `/stock/release`, `/stock/{productId}`, `/stock/movements/{productId}`)
- **Secondary ports**: Repository interfaces for PostgreSQL (stock data) and MongoDB (audit logs)
- **Strategy pattern**: Multiple reservation policies based on product type or business rules

## Technology Stack & Conventions

### Dependencies (go.mod)
- Gin framework for HTTP routing
- GORM for PostgreSQL ORM
- MongoDB Go driver for audit logging
- Kafka Go client for async processing
- Uber Zap for structured logging
- gin-swagger for API documentation
- testify/mock for testing

### Database Strategy
- **PostgreSQL**: Transactional stock operations (current balances)
- **MongoDB**: Audit trail and movement history (append-only logs)
- Use separate repositories for each data store

### Messaging Patterns
- Kafka topics for async stock operations
- Worker pool pattern for concurrent message processing
- Both sync (REST API) and async (Kafka) support for all operations

## Development Workflows

### Project Structure
```
cmd/api/           # Main application entry point
internal/domain/   # Business entities and interfaces
internal/app/      # Use cases and service layer
internal/infra/    # Database and external service adapters
internal/strategy/ # Reservation strategy implementations
internal/tests/    # Test utilities and mocks
api/docs/         # Swagger/OpenAPI documentation
scripts/          # Database migrations, seeds, Kafka setup
docker/           # Docker Compose and container configurations
```

### Testing Approach
- Unit tests for domain logic and strategies: `go test ./internal/domain/...`
- Service layer tests with mocks: `go test ./internal/app/...`
- Integration tests for repositories: `go test ./internal/infra/...`
- Use testify/mock for external dependencies

### Local Development
- Docker Compose setup with Go API, PostgreSQL, MongoDB, and Kafka
- Run complete stack: `docker compose up`
- API documentation available at: `http://localhost:8090/swagger/index.html`

## Code Conventions

### Logging
- Use Uber Zap for structured logging throughout
- Log all stock operations for audit compliance
- Include correlation IDs for tracing across services

### Error Handling
- Domain errors should be well-defined types
- Use proper HTTP status codes in API responses
- Include meaningful error messages for client consumption

### Strategy Pattern Usage
- Implement `ReservationStrategy` interface for different business rules
- Examples: `StandardReservationStrategy`, `PremiumProductStrategy`
- Strategy selection based on product type or business context

### Repository Pattern
- Define interfaces in domain layer
- Implement concrete repositories in infra layer
- Use dependency injection for testability

## API Design
- Follow RESTful conventions
- Use DTOs for request/response serialization
- Include proper HTTP status codes and error responses
- Document all endpoints with Swagger annotations

## Async Processing
- Kafka consumers should implement worker pool pattern
- Handle message processing failures gracefully
- Ensure idempotency for retry scenarios
- Use structured logging for debugging async flows

When implementing new features, maintain the separation of concerns between domain logic, application orchestration, and infrastructure concerns. Always consider both synchronous and asynchronous processing paths for stock operations.