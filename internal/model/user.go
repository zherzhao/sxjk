package model

type User struct {
	UserID   int64  `eorm:"user_id"`
	Username string `eorm:"username"`
	Password string `eorm:"password"`
}
