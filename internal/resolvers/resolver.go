package resolvers

import (
	"context"

	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/gqlgen"
	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/prisma"
)

// Resolver is the root resolver containing mutations, querys, and models.
type Resolver struct{ Prisma *prisma.Client }

// Mutation returns a MutationResolver.
func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}

// Query returns a QueryResolver.
func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *queryResolver) CheckHealth(ctx context.Context) (bool, error) {
	return true, nil
}

func (r *mutationResolver) Placeholder(ctx context.Context) (bool, error) {
	return true, nil
}
