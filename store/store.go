package store

import (
	config "app/configs"
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	DB             *sql.DB
	UserRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", config.Cfg.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *Store) Close() {
	s.DB.Close()
}

func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
	}
	return s.UserRepository
}
