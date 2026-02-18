package post

import (
	"context"
	"time"

	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *Repository {
	if conn == nil {
		logger.Fatal("connection of post repository is nil")
	}
	return &Repository{
		Conn: conn,
	}
}
func (r *Repository) Insert(post CreatePostDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.Conn.Exec(
		ctx,
		"INSERT INTO posts (username,body) VALUES ($1,$2)",
		post.Username,
		post.Body,
	)
	return err
}
