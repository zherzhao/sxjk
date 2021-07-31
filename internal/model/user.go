package model

type User struct {
	UserID   int64  `eorm:"user_id"`
	Username string `eorm:"username"`
	Password string `eorm:"password"`
	Role     string `eorm:"role"`
	Unit     string `eorm:"unit"`
}

type UserResp struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Unit     string `json:"unit"`
}

type UserReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Unit     string `json:"unit" binding:"required"`
}
