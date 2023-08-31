package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	*sqlx.DB
	*SegmentRepository
	*UserRepository
}

func NewRepository(dsn string) (*Repository, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		fmt.Println(dsn)
		log.Fatal("Cannot connect to Postgres")
	}

	return &Repository{
		DB:                db,
		SegmentRepository: &SegmentRepository{DB: db},
		UserRepository:    &UserRepository{DB: db},
	}, nil
}
