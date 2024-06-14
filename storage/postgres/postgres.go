package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Salikhov079/library-auth/config"
	u "github.com/Salikhov079/library-auth/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	Users u.User
}

func NewPostgresStorage() (u.StorageRoot, error) {
	cnf := config.Load()

	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cnf.PostgresUser, cnf.PostgresPassword,
		cnf.PostgresHost, cnf.PostgresPort,
		cnf.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, err
}

func (s *Storage) User() u.User {
	if s.Users == nil {
		s.Users = &UserStorage{s.Db}
	}
	return s.Users
}
