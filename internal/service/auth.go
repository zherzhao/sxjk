package service

import (
	"webconsole/internal/dao/database"
	"webconsole/internal/model"
	"webconsole/pkg/jwt"
	sf "webconsole/pkg/snowflake"
)

func SignUp(p *model.ParamSignUp) error {
	// 检查用户是否已经注册
	err := database.CheckUserExist(p.UserName)
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
	return database.InsertUser(user)

}

// 处理用户登录以及JWT的发放
func Login(p *model.ParamLogin) (userid int64, userRole, aToken string, err error) {
	// 构造一个User实例
	user := &model.User{
		Username: p.UserName,
		Password: p.Password,
	}

	// 数据库验证
	if err = database.UserLogin(user); err != nil {
		return
	}

	// 验证通过后发放token
	aToken, err = jwt.GenToken(user.UserID, user.Username, user.Role)
	return user.UserID, user.Role, aToken, err

}
