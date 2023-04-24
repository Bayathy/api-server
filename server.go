package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bayathy/api-server/cmd/db"
	"github.com/bayathy/api-server/graph/resolver"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bayathy/api-server/graph"
)

const defaultPort = "8080"

func main() {
	database, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	mux := http.NewServeMux()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{DB: database}}))

	handler := cors.Default().Handler(mux)
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
