package model

type ParamSignUp struct {
	VerifyCode string `json:"verifycode" binding:"required"`
	UserName   string `eorm:"username" json:"username" binding:"required"`
	Unit       string `eorm:"unit" json:"unit" binding:"required"`
	Password   string `eorm:"password" json:"password" binding:"required"`
	RePassword string `eorm:"repassword" json:"repassword" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	UserName string `eorm:"username" json:"username" binding:"required"`
	Password string `eorm:"password" json:"password" binding:"required"`
}
