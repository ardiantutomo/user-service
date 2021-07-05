package main

import (
	"log"
	"net/http"
	"os"
	"user-service/graph"
	"user-service/graph/generated"
	modeldb "user-service/model"
	"user-service/repository"
	"user-service/service"

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
	if err := userRepo.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
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
