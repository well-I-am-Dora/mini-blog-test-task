package app

import (
	"github.com/hashicorp/go-memdb"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func newPGDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func newMemDB() *memdb.MemDB {
	db, err := memdb.NewMemDB(getMemDBSchema())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getMemDBSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"posts": {
				Name: "posts",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UUIDFieldIndex{Field: "ID"},
					},
				},
			},
			"comments": {
				Name: "comments",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UUIDFieldIndex{Field: "ID"},
					},
					"post_id": {
						Name:    "post_id",
						Unique:  false,
						Indexer: &memdb.UUIDFieldIndex{Field: "PostID"},
					},
				},
			},
		},
	}
}
