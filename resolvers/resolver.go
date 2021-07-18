package resolvers

import "flashcards/generated"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	generated.ResolverRoot
}

func NewResolver() *Resolver {
	return &Resolver{}
}
