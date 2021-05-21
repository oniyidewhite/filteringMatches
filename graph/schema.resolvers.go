package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oniyidewhite/filteringMatches/graph/generated"
	"github.com/oniyidewhite/filteringMatches/graph/model"
)

func (r *queryResolver) Matches(ctx context.Context, input model.InputFilter) ([]*model.Match, error) {
	result, err := r.Repository.FindMatches(MapInputFilter(input))
	if err != nil {
		r.Logger.Println(err)
		return nil, err
	}

	return MapMatchResult(result), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
