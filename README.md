# ms-nexusMarket-inventory — Stock Microservice (MVP Cloud Ready)

[![License: Apache](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
![Docker: Build](https://img.shields.io/badge/Docker-Cloud_Build-blue.svg)
![Docs: Swagger Ui](https://img.shields.io/badge/Docs-Swagger_UI-green.svg)


## ✨ Context

**ms-nexusMarket-inventory** is a Stock Microservice built as a Minimum Viable Product (MVP) using a modern and scalable architecture. The aim is to practice and apply professional development patterns, integrating multiple technologies—messaging, concurrency, structured logging, and hybrid databases—in a scenario that closely mimics real-world distributed systems.

## 🎯 Goals

* Provide a robust service for reserving, releasing, and querying product inventory, with complete persistence and audit trail for all actions.

* Integrate modern software patterns (Hexagonal Architecture, Strategy Pattern, Service Layer, etc.) and technologies prominent in contemporary microservices.

## 🚀 Key Features

* **Reserve** product stock units

* **Release** previously reserved units

* **Query** the current stock balance by product

* **Query** product movement history (audit/log)

* Both **synchronous** API and **asynchronous** queue/Kafka operations

* Comprehensive structured logging for all operations

## 🛠️ Architecture & Patterns

* **Hexagonal (Clean) Architecture**: Clear separation between domain logic, adapters, application layer, and infrastructure

* **Strategy Pattern**: Allows flexible reservation rules by product or business type

* **Repository Pattern**, **Service Layer & DTOs**: To maximize modularity and testability

* **Worker pool (concurrency)**: Efficient parallel processing of Kafka events

## ⚙️ Tech Stack

* **Language**: Go (Golang)

* **HTTP Framework**: Gin

* **Relational Database**: PostgreSQL (with GORM)

* **NoSQL Database**: MongoDB (for logs/history)

* **Messaging**: Kafka (asynchronous inventory events)

* **Logging**: Uber Zap (structured logs for APIs and workers)

* **API Documentation**: Swagger/OpenAPI (via gin-swagger)

* **Orchestration**: Docker Compose (local stack)

* **Testing**: Go Testing, testify/mock

## 📡 RESTful Endpoints

* **POST /stock/reserve** — Reserve stock units for a chosen product

* **POST /stock/release** — Release previously reserved units

* **GET /stock/{productId}** — Query current stock for a specific product

* **GET /stock/movements/{productId}** — Query product audit/history

## 🏆 MVP Requirements

### Functional

* Multiple reservation strategies (Strategy Pattern)
* Complete audit logging on every operation
* Concurrent processing of stock actions (Kafka workers)

### Non-Functional

* Unit tests for all main use cases and strategies
* API fully documented (Swagger/OpenAPI)
* Structured logs (Zap) in all flows
* Local execution environment with Docker Compose (Go API, Postgres, Mongo, Kafka)
* Clean, portable, and easily extensible codebase

## ▶️ Getting Started

```console
# Requirements: Docker + Docker Compose installed

# Clone the repository
git clone https://github.com/Ch94Ca/ms-nexusMarket-inventory.git
cd ms-nexusMarket-inventory

# (First time only) Create the shared external Docker network:
docker network create ms-nexusmarket-network

# Start infrastructure services (Postgres, Kafka, Mongo, etc):
docker-compose -f docker-compose.infra.yaml up -d

# Start the API (run this each time you want to update code or restart just the API):
docker-compose -f docker-compose.api.yaml up --build

# Access the Swagger documentation:
# http://localhost:8090/swagger/index.html
```

## 🗂️ Project Structure

```text
    ms-nexusMarket-inventory/

    ├── api/

    │ ├── docs/

    │ └── examples/

    ├── cmd/

    │ └── api/

    ├── docker/

    │ ├── config/

    │ ├── dockerfiles/

    │ └── volumes/

    ├── internal/

    │ ├── app/

    │ ├── domain/

    │ ├── infra/

    │ ├── strategy/

    │ └── tests/

    ├── scripts/

    │ ├── kafka/

    │ ├── migrate/

    │ ├── seed/

    │ └── utils/

    ├── .gitattributes

    ├── .gitignore

    ├── LICENSE

    ├── README.md

    └── go.mod
```

### Directory Descriptions

* **api/**:
    Contains everything related to the API contracts, such as documentation (e.g., Swagger/OpenAPI files) and sample requests and responses.

* **cmd/**:
    Holds the application's entry points. Each subdirectory represents a distinct runnable service or executable, such as the main API server or background workers.

* **docker/**:
    Stores Docker-related files, including service configurations, Dockerfiles for building images, and data volumes for development or testing environments.

* **internal/**: This is the core business logic of the application.

    * **app/**: 
        Application layer containing use cases, operations, and orchestration logic.
        domain/: Domain models, entities, interfaces, and pure business rules.

    * **infra/**:
        Infrastructure implementations for things like databases, external APIs, messaging, and support libraries.

    * **strategy/**: 
        Pluggable business strategies or interchangeable components implementing domain interfaces.

    * **tests/**: 
        Helpers, mocks, and base classes for reusable testing utilities.

* **scripts/**: Auxiliary scripts to automate, seed, migrate, or maintain the system.

    * **kafka/**: 
    Scripts to create Kafka topics or set up the messaging infrastructure.
    migrate/: Database migration scripts or helpers.

    * **seed/**: 
    Scripts for populating development or demo data.
    utils/: Miscellaneous utilities, such as advanced health checks or cleanup tasks.

* **.gitattributes, .gitignore**:
    Git versioning and configuration files.

* **LICENSE**:
    The project's license terms.

* **README.md**:
    Project overview, structure, and instructions.

* **go.mod**:
    Go module definition and dependency management.

## 🧪 Testing

You’ll find unit tests covering:

* Domain business logic
* Custom reservation strategies (Strategy Pattern)
* Application and integration services

Run all tests:

```console
go test ./...
```

## 🔗 API Documentation (Swagger)

Once running, docs are available at:

http://localhost:8090/swagger/index.html

## 📄 License

Distributed under the Apache 2.0 License.

See the LICENSE file for details.

Would you like example API request/response samples or any additional usage details included?