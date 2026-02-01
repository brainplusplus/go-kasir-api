# Kasir API

Simple Point of Sale (POS) API Service built with Go, refactored to use **Hexagonal Architecture**.

## Architecture

This project follows the Hexagonal Architecture (Ports and Adapters) pattern to ensure separation of concerns and testability.

### Structure

```
├── internal/
│   ├── app.go                   # Application wiring (Dependency Injection)
│   ├── config/                  # Configuration Loading (Viper)
│   ├── domain/                  # Domain Models (Produk, Category)
│   ├── port/                    # Interfaces (Ports)
│   │   ├── inbound/             # Input Ports (Service Interfaces)
│   │   └── outbound/            # Output Ports (Repository Interfaces)
│   ├── service/                 # Service Implementations (Business Logic)
│   └── adapter/                 # Implementations (Adapters)
│       ├── inbound/             # Input Adapters (HTTP Handlers & Routes)
│       │   └── http/
│       │       ├── dto/         # Data Transfer Objects
│       │       └── handler/     # HTTP Handlers (Produk, Category)
│       └── outbound/            # Output Adapters (Repositories)
│           ├── memory/          # In-Memory Repository Implementation
│           └── persistence/     # Database Repository Implementation (GORM)
├── main.go                      # Entry point
├── .env                         # Environment Variables (Config)
└── .gitignore                   # Git Ignore Rules
```

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL (Optional, defaults to In-Memory if not configured)

### Setup

1. Copy `.env.example` to `.env`:

```env
PORT=8080
STORAGE=postgres

# Database Configuration
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=kasir_api
DB_PORT=5432
DB_SSLMODE=disable # Use 'require' for cloud DBs like Supabase
```

2. Run the application:

```bash
go run main.go
```

The server will start at `http://localhost:8080` (or the port defined in `.env`).

### SSL Mode

If you are connecting to a cloud database (e.g., Supabase, Neon) that requires SSL, likely set `DB_SSLMODE=require` in your `.env`.

**Note for Supabase Users**: Use the **Connection Pooler** (IPv4, usually port 6543) if your local network does not support IPv6.

## API Endpoints

### Products

**Response Format**: Product details now include nested Category information.

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/api/produk` | Get all products |
| `GET` | `/api/produk/{id}` | Get product by ID (includes `category` object) |
| `POST` | `/api/produk` | Create a new product (requires `category_id`) |
| `PUT` | `/api/produk/{id}` | Update a product |
| `DELETE` | `/api/produk/{id}` | Delete a product |

**Example POST Payload:**
```json
{
    "name": "Indomie",
    "price": 3500,
    "stock": 100,
    "category_id": 1
}
```

### Categories

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/api/categories` | Get all categories |
| `GET` | `/api/categories/{id}` | Get category by ID |
| `POST` | `/api/categories` | Create a new category |
| `PUT` | `/api/categories/{id}` | Update a category |
| `DELETE` | `/api/categories/{id}` | Delete a category |

### Health Check

- `GET /health` - Check API status

## CLI Helpers

- `kill_port.ps1` / `kill_port.sh` - Scripts to free up port 8080 if usage conflict occurs.
