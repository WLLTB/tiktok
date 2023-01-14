package schema

type User struct {
	UserId   int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}
