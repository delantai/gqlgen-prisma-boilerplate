package gqlgen

import (
	"context"

	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/prisma"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Placeholder(ctx context.Context) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePost(ctx context.Context, input CreatePostInput) (PostPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePost(ctx context.Context, input PostActionInput) (PostPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) PublishPost(ctx context.Context, input PostActionInput) (PostPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) Authorize(ctx context.Context, input AuthorizeInput) (AuthorizePayload, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (prisma.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) CheckHealth(ctx context.Context) (bool, error) {
	panic("not implemented")
}
func (r *queryResolver) Feed(ctx context.Context) ([]prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Drafts(ctx context.Context) ([]prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, id string) (*prisma.Post, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	panic("not implemented")
}
