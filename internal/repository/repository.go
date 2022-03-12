package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Repositories struct {
	User *UserRepo
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		User: NewUserRepo(db),
	}
}
