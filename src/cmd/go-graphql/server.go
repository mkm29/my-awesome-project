package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mkm29/my-awesome-project/src/cmd/go-graphql/graph"
	"github.com/mkm29/my-awesome-project/src/cmd/go-graphql/graph/generated"
	"github.com/mkm29/my-awesome-project/src/pkg/postgres"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	toDoService := &postgres.ToDoImpl{Message: "Hello!"}
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{ToDo: toDoService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
