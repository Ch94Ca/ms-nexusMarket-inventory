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

# Start the complete stack (API, Kafka, Postgres, Mongo)
docker compose up

# Access the Swagger documentation:
# http://localhost:8090/swagger/index.html
```

## 🗂️ Project Structure

Work In Progess

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