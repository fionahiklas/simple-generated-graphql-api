package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"simple-generated-graphql-api/graph/generated"
	"simple-generated-graphql-api/graph/model"
)

// CreateHome is the resolver for the createHome field.
func (r *mutationResolver) CreateHome(ctx context.Context, input model.NewHome) (*model.Home, error) {
	panic(fmt.Errorf("not implemented: CreateHome - createHome"))
}

// Homes is the resolver for the homes field.
func (r *queryResolver) Homes(ctx context.Context) ([]*model.Home, error) {
	panic(fmt.Errorf("not implemented: Homes - homes"))
}

// Home is the resolver for the home field.
func (r *queryResolver) Home(ctx context.Context, input string) (*model.Home, error) {
	panic(fmt.Errorf("not implemented: Home - home"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
