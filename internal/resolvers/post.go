package resolvers

import (
	"context"

	gqlgen "bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/gqlgen"
	prisma "bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/prisma"
)

// Post returns a PostResolver.
func (r *Resolver) Post() gqlgen.PostResolver {
	return &postResolver{r}
}

// Queries.
func (r *queryResolver) Feed(ctx context.Context) ([]prisma.Post, error) {
	isPublished := true
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{IsPublished: &isPublished},
	}).Exec(ctx)
}

func (r *queryResolver) Drafts(ctx context.Context) ([]prisma.Post, error) {
	isPublished := false
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{IsPublished: &isPublished},
	}).Exec(ctx)
}

func (r *queryResolver) Post(
	ctx context.Context,
	id string,
) (*prisma.Post, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &id}).Exec(ctx)
}

// Mutations.
func (r *mutationResolver) CreatePost(
	ctx context.Context,
	input gqlgen.CreatePostInput,
) (gqlgen.PostPayload, error) {
	post, err := r.Prisma.CreatePost(prisma.PostCreateInput{
		Author: prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{ID: &input.AuthorID},
		},
		Content: input.Content,
		Title:   input.Title,
	}).Exec(ctx)
	return gqlgen.PostPayload{Post: post}, err
}

func (r *mutationResolver) DeletePost(
	ctx context.Context,
	input gqlgen.PostActionInput,
) (gqlgen.PostPayload, error) {
	post, err := r.Prisma.DeletePost(
		prisma.PostWhereUniqueInput{ID: &input.ID},
	).Exec(ctx)
	return gqlgen.PostPayload{Post: post}, err
}

func (r *mutationResolver) PublishPost(
	ctx context.Context,
	input gqlgen.PostActionInput,
) (gqlgen.PostPayload, error) {
	isPublished := true
	post, err := r.Prisma.UpdatePost(prisma.PostUpdateParams{
		Data:  prisma.PostUpdateInput{IsPublished: &isPublished},
		Where: prisma.PostWhereUniqueInput{ID: &input.ID},
	}).Exec(ctx)
	return gqlgen.PostPayload{Post: post}, err
}

// Posts.
type postResolver struct{ *Resolver }

func (r *postResolver) Author(
	ctx context.Context,
	obj *prisma.Post,
) (prisma.User, error) {
	author, err := r.Prisma.Post(prisma.PostWhereUniqueInput{
		ID: &obj.ID,
	}).Author().Exec(ctx)
	return *author, err
}
