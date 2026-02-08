package internal

import (
	"encoding/json"
	"fmt"
	httpAdapter "kasir-api/internal/adapter/inbound/http"
	httpHandler "kasir-api/internal/adapter/inbound/http/handler"
	"kasir-api/internal/adapter/outbound/memory"
	"kasir-api/internal/adapter/outbound/persistence"
	"kasir-api/internal/adapter/outbound/persistence/entity"
	"kasir-api/internal/config"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/port/outbound"
	"kasir-api/internal/service"
	"kasir-api/internal/utils/swagger"
	"net/http"

	"gorm.io/gorm"
)

func NewApp(cfg *config.Config, db *gorm.DB) http.Handler {
	var produkRepo outbound.ProdukRepository
	var categoryRepo outbound.CategoryRepository
	var transactionRepo outbound.TransactionRepository
	var reportRepo outbound.ReportRepository
	var tm outbound.TransactionManager

	// Initialize Repository based on Config
	if cfg.Storage == "postgres" && db != nil {
		fmt.Println("Using Postgres Storage")
		produkRepo = persistence.NewDBProdukRepository(db)
		categoryRepo = persistence.NewDBCategoryRepository(db)
		transactionRepo = persistence.NewDBTransactionRepository(db)
		reportRepo = persistence.NewDBReportRepository(db)
		tm = persistence.NewGormTransactionManager(db)

		// Auto Migrate Transactions
		if err := db.AutoMigrate(&entity.Transaction{}, &entity.TransactionDetail{}); err != nil {
			fmt.Printf("Failed to auto migrate transactions: %v\n", err)
		}
	} else {
		fmt.Println("Using In-Memory Storage")
		produkRepo = memory.NewInMemoryProdukRepository()
		categoryRepo = memory.NewInMemoryCategoryRepository()
		// In-memory transaction repo not implemented yet, panic or ignore
		// For now, let's assume postgres is used as per requirement
	}

	// Initialize Services
	produkService := service.NewProdukService(produkRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	var transactionService inbound.TransactionService
	if transactionRepo != nil && reportRepo != nil && tm != nil {
		transactionService = service.NewTransactionService(transactionRepo, produkRepo, reportRepo, tm)
	}

	// Initialize Swagger Generator
	swag := swagger.NewGenerator("Kasir API", "1.0.0")

	// Initialize Handler with Swagger
	produkHandler := httpHandler.NewProdukHandler(produkService, swag)
	categoryHandler := httpHandler.NewCategoryHandler(categoryService, swag)
	var transactionHandler *httpHandler.TransactionHandler
	if transactionService != nil {
		transactionHandler = httpHandler.NewTransactionHandler(transactionService, swag)
	}

	// Initialize Router
	router := httpAdapter.NewRouter(produkHandler, categoryHandler)

	mux := http.NewServeMux()

	// Register routes
	router.RegisterRoutes(mux)
	if transactionHandler != nil {
		transactionHandler.RegisterRoutes(mux)
	}

	// Swagger JSON
	mux.Handle("/swagger.json", swag)

	// Swagger UI
	mux.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>SwaggerUI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
<script>
    window.onload = () => {
        window.ui = SwaggerUIBundle({
            url: '/swagger.json',
            dom_id: '#swagger-ui',
        });
    };
</script>
</body>
</html>
		`))
	})

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	return httpAdapter.CORSMiddleware(mux)
}
