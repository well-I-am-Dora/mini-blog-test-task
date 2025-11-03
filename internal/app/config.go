package app

import "os"

type config struct {
	inMemory bool
}

func newConfig() *config {
	return &config{
		inMemory: os.Getenv("IN_MEMORY") == "true",
	}
}
