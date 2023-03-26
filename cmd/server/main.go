package main

import (
	"context"
	"log"

	"github.com/bayathy/api-server/ent"
)

func main() {
    client, err := ent.Open("postgres","host=db port=5432 user=root dbname=postgres password=root sslmode=disable")
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

}