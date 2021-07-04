package main

import (
	"go-graphql-clean/graph"
	"go-graphql-clean/graph/generated"
	modeldb "go-graphql-clean/model"
	"go-graphql-clean/repository"
	"go-graphql-clean/service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, _ := modeldb.DBConnection()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userResolver := graph.NewUserResolver(userService)
	userHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: userResolver,
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", userHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
