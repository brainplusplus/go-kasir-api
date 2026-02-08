# Architecture Documentation

This project follows the **Hexagonal Architecture** (also known as Ports and Adapters) to ensure a clear separation of concerns, testability, and maintainability.

## 1. High-Level Overview

The core business logic (Domain & Services) is isolated from external technologies (Database, HTTP, Config). This allows us to change frameworks or databases with minimal impact on the core logic.

```mermaid
graph TD
    Client[HTTP Client] --> AdapterIn[Input Adapters (HTTP Handlers)]
    AdapterIn --> PortIn[Input Ports (Service Interfaces)]
    PortIn --> Service[Service Implementation (Business Logic)]
    Service --> PortOut[Output Ports (Repository Interfaces)]
    PortOut --> AdapterOut[Output Adapters (Persistence/Memory)]
    AdapterOut --> DB[(Database)]
```

## 2. Directory Structure & Responsibilities

### `internal/domain`
Contains the **Core Entities** of the application. These are pure Go structs with no logic dependencies.
- `Category`, `Produk`, `Transaction`, `Report`
- Structs may contain JSON tags for API serialization but should remain framework-agnostic otherwise.

### `internal/port`
Defines the **Contracts** (Interfaces) for the application.
- **Inbound Ports (`internal/port/inbound`)**: Defines methods *available to* outside actors (e.g., HTTP handlers). Implemented by Services.
    - `TransactionService`, `CategoryService`, `ProdukService`
- **Outbound Ports (`internal/port/outbound`)**: Defines methods *required by* the application to store/retrieve data. Implemented by Repositories.
    - `TransactionRepository`, `ReportRepository`, `CategoryRepository`

### `internal/service`
Implements the **Inbound Ports**. This is where the **Business Logic** resides.
- Validates inputs.
- Orchestrates data flow.
- Manages Transactions (using `TransactionManager`).
- Calls Outbound Ports (Repositories).
- **Standard**: All methods accept `ctx context.Context` as the first argument.

### `internal/adapter`
Contains the implementations of interactions with the outside world.
- **Inbound Adapters (`internal/adapter/inbound/http`)**:
    - **Handlers**: Parse HTTP requests, call Services, and write HTTP responses.
    - **DTOs**: Data Transfer Objects to decouple Domain from API contracts.
- **Outbound Adapters (`internal/adapter/outbound`)**:
    - **Persistence**: Implementation of Repositories using **GORM**.
    - **Memory**: In-memory reference implementation (for testing/dev).

### `internal/config`
Handles configuration loading using **Viper**.
- Loads from `.env` or system environment variables.
- Provides a centralized configuration struct.

### `internal/app.go`
The **Wiring Center**.
- Initializes Repositories.
- Initializes Services.
- Injects dependencies (Dependency Injection).
- Sets up the HTTP Router.

## 3. The "Gold Standard" Patterns

We adhere to strict patterns to ensure consistency and reliability:

### Context Propagation
*   **All Layers**: Every method in the call stack (Handler -> Service -> Repository) accepts `ctx context.Context` as its first argument.
*   **Purpose**: Enables cancellation, timeout propagation, and request tracing/logging.
*   **Usage**: handlers pass `r.Context()`; services pass it to repositories; repositories use `db.WithContext(ctx)`.

### Repository Pattern
*   **Interfaces**: Defined in `internal/port/outbound`.
*   **Context-Aware**: All methods accept `context.Context`.
*   **Transaction-Ready**: Implementations use a helper `getDB(ctx, standardDB)` to automatically participate in active transactions found in the context.

### Transaction Management
*   **Service Layer**: Transactions are managed in the Service Layer, where the unit of work is defined.
*   **`RunInTransaction`**: A helper method injects the transaction into the context. Services explicitly wrap atomic operations in this block.
    ```go
    // Example
    err := s.tm.RunInTransaction(ctx, func(ctx context.Context) error {
        // All repo calls here automatically use the transaction in ctx
        repo.Create(ctx, data)
        return nil
    })
    ```

### Clean DTOs
*   **Request/Response Separation**: We do not bind API requests directly to Domain entities.
*   **Mapping**: Explicit mapping functions (e.g., `FromDomain`) convert between Domain entities and DTOs.
*   **Validation**: Handlers validate DTOs before passing data to Services.

## 4. Technology Stack

*   **Language**: Go 1.24+
*   **Framework**: Standard `net/http` with `http.ServeMux` (Go 1.22+ routing).
*   **Database**: PostgreSQL / SQLite (In-Memory).
*   **ORM**: GORM (v2).
*   **Config**: Viper.
*   **Documentation**: Swagger (generated/manual).

## 5. Future Considerations
*   **Authentication**: Implementing JWT middleware in `adapter/inbound/http/middleware`.
*   **Logging**: Structured logging (slog/zap) injected into context.
*   **Testing**: Adding Mocks for Ports using `mockgen` for unit testing Services.
