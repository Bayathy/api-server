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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{DB: database}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	handler := cors.Default().Handler(srv) // ★CORS レスポンス対応
	http.Handle("/query", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
