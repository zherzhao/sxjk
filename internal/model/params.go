package model

type ParamSignUp struct {
	UserName   string `eorm:"username" json:"username" binding:"required"`
	Password   string `eorm:"password" json:"password" binding:"required"`
	RePassword string `eorm:"repassword" json:"repassword" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	UserName string `eorm:"username" json:"username" binding:"required"`
	Password string `eorm:"password" json:"password" binding:"required"`
}
