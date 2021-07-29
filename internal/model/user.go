package model

type User struct {
	UserID   int64  `eorm:"user_id"`
	Username string `eorm:"username"`
	Password string `eorm:"password"`
	Role     string `eorm:"role"`
	Unit     string `eorm:"unit"`
}

type RespUser struct {
	UserID    int64  `json:"user_id"`
	UserIDStr string `json:"user_id_str"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Unit      string `json:"unit"`
}
