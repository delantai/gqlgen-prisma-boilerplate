package resolvers

import (
	"context"

	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/gqlgen"
	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/prisma"
)

// User returns a UserResolver.
func (r *Resolver) User() gqlgen.UserResolver {
	return &userResolver{r}
}

// Queries.
// Mutations.
func (r *mutationResolver) Authorize(
	ctx context.Context,
	input gqlgen.AuthorizeInput,
) (gqlgen.AuthorizePayload, error) {
	user, err := r.Prisma.CreateUser(prisma.UserCreateInput{
		Email: input.Email,
		Name:  input.Name,
	}).Exec(ctx)
	return gqlgen.AuthorizePayload{User: user}, err
}

// Users.
type userResolver struct{ *Resolver }

func (r *userResolver) Posts(
	ctx context.Context,
	obj *prisma.User,
) ([]prisma.Post, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).Posts(nil).Exec(ctx)
}
