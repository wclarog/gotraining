package main

import (
	"context"
	"log"
	"os"

	"excercise-library/ent"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	connectionString := "root:RooT27134668@tcp(localhost:3306)/dev_library"
	client, errOpen := ent.Open("mysql", connectionString)
	if errOpen != nil {
		panic("database connection failed")
	}
	defer client.Close()

	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
