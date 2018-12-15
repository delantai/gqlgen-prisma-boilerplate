package middleware

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/vektah/gqlparser/gqlerror"
)

// ErrorPresenter returns options to be passed into the GraphQL handler.
func ErrorPresenter() handler.Option {
	return handler.ErrorPresenter(errorPresenter)
}

// PanicRecoverer returns options to be passed into the GraphQL handler.
func PanicRecoverer() handler.Option {
	return handler.RecoverFunc(panicRecoverer)
}

func errorPresenter(ctx context.Context, e error) *gqlerror.Error {
	return graphql.DefaultErrorPresenter(ctx, e)
}

func panicRecoverer(ctx context.Context, err interface{}) error {
	// Notify bug tracker...
	fmt.Printf("did I get err in panic? %v\n", err)
	return fmt.Errorf("internal server error")
}
