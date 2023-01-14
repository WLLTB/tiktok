package schema

type User struct {
	UserID   int    `gorm:"primary_key;column:user_id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}
