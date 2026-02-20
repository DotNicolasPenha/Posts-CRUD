package post

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
