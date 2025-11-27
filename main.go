package main

import (
	"context"
	"fmt"
	"log"

	"github.com/itsDrac/godo/internal/db"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://todo_user:todo_pass@localhost:5432/todo_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	_ = db.New(conn)
	fmt.Println("Complete rest.")
	fmt.Println("Create a web server")
	fmt.Println("Add routing")
	fmt.Println("Add service functions in routing.")
	fmt.Println("In service struct create an argument of type Querier and pass db.New(conn) to it.")
}