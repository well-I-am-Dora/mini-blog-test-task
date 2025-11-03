package api

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	postService    postService
	commentService commentService
}

func NewResolver(
	postService postService,
	commentService commentService,
) *Resolver {
	return &Resolver{
		postService:    postService,
		commentService: commentService,
	}
}
