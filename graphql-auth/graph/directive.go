package graph

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

type Token struct{}

var Directive = DirectiveRoot{
	Auth: Authorication,
}

func Authorication(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	log.Print(ctx.Value(Token{}))
	token := ctx.Value(Token{})
	if v, ok := token.([]string); ok && len(v) == 1 && v[0] == "test-token" {
		return next(ctx)
	}
	return nil, errors.New("Not Authorized")
}
