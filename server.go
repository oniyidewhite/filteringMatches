package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/oniyidewhite/filteringMatches/graph"
	"github.com/oniyidewhite/filteringMatches/graph/generated"
	"github.com/oniyidewhite/filteringMatches/store"
	"io/ioutil"
	"log"
	"os"
)

const GraphQLEndpoint = "/graphql"
const PORT = "8080"
const AppShortName = "ptest"
const Matchesfile = "matches.json"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = PORT
	}

	logger := NewLogger()
	repo, err := setupRepository(Matchesfile)
	if err != nil {
		logger.Fatalln(err)
	}

	engine := gin.Default()
	engine.GET("/", NewPlayground())
	engine.POST(GraphQLEndpoint, NewGraphQLServer(logger, repo))
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		logger.Fatalln(err)
	}
}

func setupRepository(fileUrl string) (*store.Repository, error) {
	fileData, err := ioutil.ReadFile(fileUrl)
	if err != nil {
		return nil, err
	}

	repo, err := store.NewRepository(fileData)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func NewLogger() *log.Logger {
	return log.New(os.Stdout, AppShortName, log.LstdFlags|log.Lshortfile)
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
