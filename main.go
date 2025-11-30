package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/itsDrac/godo/handler"
	"github.com/itsDrac/godo/internal/db"
	"github.com/itsDrac/godo/internal/service"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://todo_user:todo_pass@localhost:5432/todo_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	query := db.New(conn)
	userService := service.NewUserService(query)
	
	handler := handler.NewChiHandler(userService)
	handler.Mount()
	// Server needs handler.
	serv := http.Server{
		Addr: ":8080",
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