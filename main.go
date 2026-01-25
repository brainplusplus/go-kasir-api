package main

import (
	"fmt"
	"kasir-api/internal"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := internal.NewApp()

	fmt.Printf("Server running di localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
