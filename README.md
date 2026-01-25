# Kasir API

Simple Point of Sale (POS) API Service built with Go, refactored to use **Hexagonal Architecture**.

## Architecture

This project follows the Hexagonal Architecture (Ports and Adapters) pattern to ensure separation of concerns and testability.

### Structure

```
├── internal/
│   ├── app.go                   # Application wiring (Dependency Injection)
│   ├── service/                 # Service Implementations (Business Logic)
│   │   ├── category.go          # Category Business Logic
│   │   └── produk.go            # Product Business Logic
│   ├── port/                    # Interfaces (Ports)
│   │   ├── inbound/             # Input Ports (Service Interfaces)
│   │   └── outbound/            # Output Ports (Repository Interfaces)
│   ├── adapter/                 # Implementations (Adapters)
│   │   ├── inbound/             # Input Adapters (HTTP Handlers & Routes)
│   │   │   ├── http/            # HTTP Adapter
│   │   │   │   ├── handler/     # HTTP Handlers (Produk, Category)
│   │   │   │   ├── middleware.go # HTTP Middlewares
│   │   │   │   └── routes.go    # Router Configuration
│   │   └── outbound/            # Output Adapters (Repositories)
│   │       └── memory/          # In-Memory Repository Implementation
│   ├── model/                   # Data Models (Structs)
│   └── utils/                   # Utilities
│       └── swagger/             # Custom Swagger Generator
├── main.go                      # Entry point
├── kill_port.ps1                # Helper to kill app port (Windows PowerShell)
├── kill_port.sh                 # Helper to kill app port (Windows Git Bash)
├── kill_port_linux.sh           # Helper to kill app port (Linux Native)
├── .env                         # Environment Variables (Config)
└── .gitignore                   # Git Ignore Rules
```

## Getting Started

### Prerequisites

- Go 1.20+

### Setup

1. Copy `.env.example` to `.env` (or create one):

```env
PORT=6001
```

2. Run the application:

```bash
go run main.go
```

The server will start at `http://localhost:6001` (or the port defined in `.env`).

### Helper Scripts

If you encounter "Port already in use" errors:

- **Windows (PowerShell)**: `.\kill_port.ps1`
- **Windows (Git Bash)**: `./kill_port.sh`
- **Linux**: `./kill_port_linux.sh`

## API Documentation (Swagger)

The API documentation is automatically generated.

- **Swagger UI**: [http://localhost:6001/swagger](http://localhost:6001/swagger)
- **JSON Spec**: [http://localhost:6001/swagger.json](http://localhost:6001/swagger.json)

## API Endpoints

### Products

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/api/produk` | Get all products |
| `GET` | `/api/produk/{id}` | Get product by ID |
| `POST` | `/api/produk` | Create a new product |
| `PUT` | `/api/produk/{id}` | Update a product |
| `DELETE` | `/api/produk/{id}` | Delete a product |

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
