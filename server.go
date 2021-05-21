package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/oblessing/filteringMatches/graph"
	"github.com/oblessing/filteringMatches/graph/generated"
	"github.com/oblessing/filteringMatches/store"
	"log"
	"os"
)

const GraphQLEndpoint = "/graphql"
const PORT = "8080"
const AppShortName = "ptest"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port := os.Getenv("PORT")
		if port == "" {
			port = PORT
		}

		logger := NewLogger()
		repo := store.NewRepository()

		engine := gin.Default()
		engine.GET("/", NewPlayground())
		engine.POST(GraphQLEndpoint, NewGraphQLServer(logger, repo))
		if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
			logger.Fatalln(err)
		}
	}
}


func NewLogger() *log.Logger {
	return log.New(os.Stdout, AppShortName, log.LstdFlags | log.Lshortfile)
}

func NewGraphQLServer(logger *log.Logger, repository *store.Repository) gin.HandlerFunc {
	config := generated.Config{Resolvers: &graph.Resolver{
		Logger:     logger,
		Repository: repository,
	}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(context *gin.Context) {
		srv.ServeHTTP(context.Writer, context.Request)
	}
}

func NewPlayground() gin.HandlerFunc {
	srv := playground.Handler("GraphQL playground", GraphQLEndpoint)
	return func(context *gin.Context) {
		srv.ServeHTTP(context.Writer, context.Request)
	}
}

