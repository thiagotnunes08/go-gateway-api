package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/thiagotnunes08/go-gateway-api/internal/repository"
	"github.com/thiagotnunes08/go-gateway-api/internal/service"
	"github.com/thiagotnunes08/go-gateway-api/internal/web/server"
)

func main() {
	if err := godotenv.Load(); err != nil {

		log.Fatal("Error loading .env file")
	}

	connsStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s name=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := sql.Open("postgres", connsStr)

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(accountService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
