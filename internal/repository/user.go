package repository

import (
	"architecture/internal/core"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db,
	}
}

func (repo *UserRepo) Get(ctx context.Context, id int) core.User {
	var user core.User

	err := repo.db.QueryRow(ctx, "SELECT id, name, email FROM person WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		fmt.Print("User repo query error", err)
	}

	return user
}
