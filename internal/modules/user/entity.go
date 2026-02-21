package post

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}
