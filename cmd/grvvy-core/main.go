package main

import (
	"bitbucket.org/delantai/gqlgen-prisma-boilerplate/internal/server"
)

func main() {
	s := server.New()
	s.Run()
}
