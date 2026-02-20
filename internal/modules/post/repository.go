package post

import (
	"context"
	"time"

	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
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
func (r *Repository) FindMany() ([]Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var posts []Post
	rows, err := r.Conn.Query(ctx, "SELECT ID,username,body,created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID
		var nameuser string
		var body string
		var createdAt time.Time

		if err := rows.Scan(&id, &nameuser, &body, &createdAt); err != nil {
			return nil, err
		}

		post := Post{
			ID:        id,
			Username:  nameuser,
			Body:      body,
			CreatedAt: createdAt,
		}

		posts = append(posts, post)
	}
	return posts, nil
}
