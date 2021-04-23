package v1

import (
	"webconsole/internal/model"
	"webconsole/pkg/jwt"
)

func SignUp(p *model.ParamSignUp) error {
	// 检查用户是否已经注册
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		return err
	}

	// 生成UID
	userID := sf.GenID()
	// 构造一个User实例
	user := &model.User{
		UserID:   userID,
		Username: p.UserName,
		Password: p.Password,
	}

	// 存入数据库
	return mysql.InsertUser(user)

}

// 处理用户登录以及JWT的发放
func Login(p *models.ParamLogin) (aoken string, err error) {
	// 构造一个User实例
	user := &models.User{
		Username: p.UserName,
		Password: p.Password,
	}

	// 数据库验证
	if err = mysql.UserLogin(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)

}
