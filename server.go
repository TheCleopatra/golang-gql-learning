package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TheCleopatra/golang-gql-learning/http"
	"github.com/TheCleopatra/golang-gql-learning/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	server := gin.Default()

	server.Use(middleware.BasicAuth())

	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", http.GraphQLHandler())

	server.Run(":" + port)
}
