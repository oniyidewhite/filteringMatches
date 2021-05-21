package graph

import (
	"github.com/oblessing/filteringMatches/store"
	"log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Logger *log.Logger
	Repository *store.Repository
}
