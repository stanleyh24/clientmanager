package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(s *pgxpool.Pool) *Store {
	return &Store{db: s}
}
