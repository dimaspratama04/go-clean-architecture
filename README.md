# Go Clean Architecture

**A Clean Architecture implementation in Go**

## 📌 Overview

This project is a backend boilerplate written in **Go**, following the **Clean Architecture** pattern.  
It separates the application layers (delivery, use case, repository, and infrastructure) to make the codebase modular, testable, and maintainable.

## 🛠 Tech Stack

- **Golang** (v1.24+)
- **GoFiber v2.52** – Web framework
- **MySQL 8** – Relational database
- **Redis 8.2.1** – Caching and session storage
- **RabbitMQ 4.1.4** – Message queue / event bus
- **Docker Compose** – Container orchestration for local development

### Project Structure

```
/cmd                → Application entry point
/db/migrations      → Database migration scripts
/internal           → Core business logic, repositories, and handlers
.env.example        → Environment variable sample file
docker-compose.yaml → Local infrastructure configuration
```

## 🎯 Goals

- Keep **business logic independent** from frameworks and external dependencies.
- Make the code **testable, scalable, and maintainable**.
- Allow **easy replacement** of infrastructure (e.g., switch from MySQL to PostgreSQL).
- Serve as a **template** for new Go backend services.

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/dimaspratama04/go-clean-architecture.git
cd go-clean-architecture
```

### 2. Set up environment variables

```bash
cp .env.example .env
```

Edit the `.env` file and fill in configuration values (database, Redis, RabbitMQ, etc.).

### 3. Start dependencies

```bash
docker-compose up -d
```

### 4. Run database migrations

Execute migration scripts under `db/migrations` (depending on your migration tool).

### 5. Run the application

```bash
go run ./cmd/main.go
```

Then open your browser or API client and access your endpoints at:  
`http://localhost:8080`

## 🧭 Architecture Overview

- **Domain / Entity** – Core business models, independent of frameworks.
- **Use Case (Service)** – Application logic orchestrating repositories and entities.
- **Repository Interface** – Abstract layer defining how data is accessed.
- **Infrastructure Layer** – Concrete implementation (MySQL, Redis, RabbitMQ).
- **Delivery Layer** – HTTP handlers or controllers.
- **cmd/** – Main entry point that wires dependencies and starts the server.
- **db/migrations/** – SQL migration scripts for schema changes.

## ✅ Features

- Clean Architecture folder structure
- MySQL database integration
- Redis caching
- RabbitMQ message queue integration
- Environment-based configuration
- Ready-to-use Docker Compose setup
- Modular and easily extendable code

## 🧪 Example API Endpoint

You can create a simple example like this:

**GET /api/v1/healthz**

```json
{
  "status": "ok",
  "message": "service is running"
}
```

**POST /api/products**

```json
{
  "product_name": "Laptop",
  "price": 1200,
  "category": "Electronics"
}
```

Response:

```json
{
  "status": "ok",
  "message": "product successfully created",
  "data": {
    "id": 1,
    "product_name": "Laptop",
    "price": 1200,
    "category": "Electronics"
  }
}
```
