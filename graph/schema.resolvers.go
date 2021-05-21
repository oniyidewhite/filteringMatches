package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/oblessing/filteringMatches/graph/generated"
	"github.com/oblessing/filteringMatches/graph/model"
)

func (r *queryResolver) Matches(ctx context.Context, input model.Filter) ([]*model.Match, error) {
	r.Logger.Println(fmt.Errorf("not implemented"))

	return nil, fmt.Errorf("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
