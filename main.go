package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/itsDrac/godo/handler"
	"github.com/itsDrac/godo/internal/db"
	"github.com/itsDrac/godo/internal/service"
	"github.com/itsDrac/godo/internal/tokens"
	"github.com/itsDrac/godo/utils"

	"github.com/jackc/pgx/v5"
)

func runMigrations(ctx context.Context, conn *pgx.Conn) error {
	migrationFiles, err := filepath.Glob("migrations/*.up.sql")
	if err != nil {
		return fmt.Errorf("failed to find migration files: %w", err)
	}
	sort.Strings(migrationFiles)

	for _, file := range migrationFiles {
		log.Printf("Running migration: %s\n", file)
		sql, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		_, err = conn.Exec(ctx, string(sql))
		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("  (table/schema already exists, skipping)\n")
			} else {
				return fmt.Errorf("failed to execute migration %s: %w", file, err)
			}
		} else {
			log.Printf("  Migration applied successfully\n")
		}
	}
	return nil
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://todo_user:todo_pass@localhost:5432/todo_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)
	if err := runMigrations(ctx, conn); err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}
	passwordService := service.NewPasswordService(
		utils.GetEnvAsInt("HASH_COST", 12),
	)

	query := db.New(conn)
	userService := service.NewUserService(query, passwordService)

	tokenizer := tokens.NewJWTTokenizer(
		utils.GetEnv("JWT_SECRET", "dev-secret"),
		24*time.Hour,
	)
	authService := service.NewAuthService(
		userService,
		passwordService,
		tokenizer,
	)
	handler := handler.NewChiHandler(userService, authService)
	handler.Mount()
	// Server needs handler.
	serv := http.Server{
		Addr:    ":8080",
		Handler: handler.Router(),
	}
	go func() {
		fmt.Printf("Starting server on %s\n", serv.Addr)
		if err := serv.ListenAndServe(); err != nil {
			log.Fatalf("Unable to start server: %v\n", err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Closed signal received")
	fmt.Println("Closing database connection.")
	conn.Close(ctx)
	fmt.Println("Shutting down server...")
	serv.Shutdown(ctx)

	fmt.Println("Add service functions in routing.")
	fmt.Println("In service struct create an argument of type Querier and pass db.New(conn) to it.")

}
