package post

type CreatePostDTO struct {
	Username string `json:"username"`
	Body     string `json:"body"`
}
