package main

import (
	"flashcards/generated"
	"flashcards/resolvers"

	"github.com/apex/gateway"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"log"

	"os"
)

func main() {

	defaultPort := ":3000"
	port := os.Getenv("PORT")
	if port == "" || port == ":" {
		port = defaultPort
	}

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// graphql
	server.GET("/__", HandleGraphqlPlayground())
	server.POST("/query", HandleGraphqlQuery())

	log.Fatal(gateway.ListenAndServe(port, nil))
}

func HandleGraphqlPlayground() gin.HandlerFunc {
	h := playground.Handler("Graphql playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func HandleGraphqlQuery() gin.HandlerFunc {
	config := generated.Config{Resolvers: resolvers.NewResolver()}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
