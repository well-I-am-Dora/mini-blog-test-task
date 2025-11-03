package app

import (
	"github.com/hashicorp/go-memdb"
	"github.com/jmoiron/sqlx"
	"github.com/martyushova/posts/internal/api"
	"github.com/martyushova/posts/internal/repository/comment_in_memory"
	"github.com/martyushova/posts/internal/repository/comment_pg"
	"github.com/martyushova/posts/internal/repository/post_in_memory"
	"github.com/martyushova/posts/internal/repository/post_pg"
	"github.com/martyushova/posts/internal/service/comment"
	"github.com/martyushova/posts/internal/service/post"
)

type Container struct {
	// Resolver
	resolver *api.Resolver
	// Service
	postService    *post.Service
	commentService *comment.Service
	// Repository
	postInMemoryRepository    *post_in_memory.Repository
	postPGRepository          *post_pg.Repository
	commentInMemoryRepository *comment_in_memory.Repository
	commentPGRepository       *comment_pg.Repository
	// Other
	conf  *config
	db    *sqlx.DB
	memDb *memdb.MemDB
}

func NewContainer() *Container {
	c := &Container{
		conf: newConfig(),
	}
	if c.conf.inMemory {
		c.memDb = newMemDB()
	} else {
		c.db = newPGDB()
	}

	return c
}

func (c *Container) GetResolver() *api.Resolver {
	if c.resolver == nil {
		c.resolver = api.NewResolver(
			c.getPostService(),
			c.getCommentService(),
		)
	}

	return c.resolver
}

// Service

func (c *Container) getPostService() *post.Service {
	if c.postService == nil {
		if c.conf.inMemory {
			c.postService = post.NewService(
				c.getPostInMemoryRepository(),
			)
		} else {
			c.postService = post.NewService(
				c.getPostPGRepository(),
			)
		}
	}

	return c.postService
}

func (c *Container) getCommentService() *comment.Service {
	if c.commentService == nil {
		if c.conf.inMemory {
			c.commentService = comment.NewService(
				c.getCommentInMemoryRepository(),
			)
		} else {
			c.commentService = comment.NewService(
				c.getCommentPGRepository(),
			)
		}
	}

	return c.commentService
}

// Repository

func (c *Container) getPostInMemoryRepository() *post_in_memory.Repository {
	if c.postInMemoryRepository == nil {
		c.postInMemoryRepository = post_in_memory.NewRepository(c.memDb)
	}
	return c.postInMemoryRepository
}

func (c *Container) getPostPGRepository() *post_pg.Repository {
	if c.postPGRepository == nil {
		c.postPGRepository = post_pg.NewRepository(c.db)
	}
	return c.postPGRepository
}

func (c *Container) getCommentInMemoryRepository() *comment_in_memory.Repository {
	if c.commentInMemoryRepository == nil {
		c.commentInMemoryRepository = comment_in_memory.NewRepository(c.memDb)
	}
	return c.commentInMemoryRepository
}

func (c *Container) getCommentPGRepository() *comment_pg.Repository {
	if c.commentPGRepository == nil {
		c.commentPGRepository = comment_pg.NewRepository(c.db)
	}
	return c.commentPGRepository
}
