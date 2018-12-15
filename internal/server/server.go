package server

import (
	"log"
	"net/http"
	"os"

	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/gqlgen"
	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/build/prisma"
	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/middleware"
	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/resolvers"
	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "4000"

// Server holds server configuration and state.
type Server struct{}

// New returns an instance of Server with any injected configuration.
func New() *Server {
	return &Server{}
}

// Run executes the server.
func (s *Server) Run() error {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	client := prisma.New(nil)
	resolver := resolvers.Resolver{
		Prisma: client,
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	http.Handle("/query", handler.GraphQL(
		gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &resolver}),
		middleware.ErrorPresenter(),
		middleware.PanicRecoverer(),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, nil)
}
