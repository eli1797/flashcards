package main

import (
	"flashcards/graph"
	"flashcards/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":3000"

func main() {

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// graphql
	server.GET("/__", HandleGraphqlPlayground())
	server.POST("/query", HandleGraphqlQuery())

	server.Run(defaultPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func HandleGraphqlPlayground() gin.HandlerFunc {
	h := playground.Handler("Graphql playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func HandleGraphqlQuery() gin.HandlerFunc {
	config := generated.Config{Resolvers: &graph.Resolver{}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
