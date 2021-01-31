package entity

type PostDetails struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
