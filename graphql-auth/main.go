package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"playground/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers:  &graph.Resolver{},
			Directives: graph.Directive,
		}))

	middleware := func(handler http.Handler) http.Handler {
		handle := func(w http.ResponseWriter, r *http.Request) {
			token := r.Header["Authorization"]
			ctx := context.WithValue(r.Context(), graph.Token{}, token)
			handler.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(handle)
	}
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
	http.Handle("/query", middleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
