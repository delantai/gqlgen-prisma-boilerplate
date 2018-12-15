// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

import (
	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/prisma"
)

type AuthorizeInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type AuthorizePayload struct {
	User *prisma.User `json:"user"`
}

type CreatePostInput struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"authorID"`
}

type PostActionInput struct {
	ID string `json:"id"`
}

type PostPayload struct {
	Post *prisma.Post `json:"post"`
}
