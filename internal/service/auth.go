package service

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"webconsole/global"
	"webconsole/internal/dao/database"
	"webconsole/internal/model"
	"webconsole/pkg/jwt"
	sf "webconsole/pkg/snowflake"
)

var userMap = map[string]struct{}{
	"太原": struct{}{},
	"大同": struct{}{},
	"阳泉": struct{}{},
	"长治": struct{}{},
	"晋城": struct{}{},
	"朔州": struct{}{},
	"晋中": struct{}{},
	"运城": struct{}{},
	"忻州": struct{}{},
	"临汾": struct{}{},
	"吕梁": struct{}{},
}

func SignUp(p *model.ParamSignUp, noUser bool) error {
	// 检查用户是否已经注册
	err := database.CheckUserExist(p.UserName)
	if err != nil {
		return err
	}
	// 生成UID
	userID := sf.GenID()

	// 构造一个User实例
	user := new(model.User)
	if noUser {
		user = &model.User{
			UserID:   userID,
			Username: p.UserName,
			Password: p.Password,
			Unit:     "交科",
			Role:     "root",
		}
	} else {
		if _, ok := userMap[p.Unit]; !ok {
			return database.ErrorInvalidUnit
		}
		user = &model.User{
			UserID:   userID,
			Username: p.UserName,
			Password: p.Password,
			Unit:     p.Unit,
			Role:     "user",
		}
	}

	// 存入数据库
	return database.InsertUser(user)

}

// 处理用户登录以及JWT的发放
func Login(p *model.ParamLogin) (int64, string, string, string, error) {
	// 构造一个User实例
	user := &model.User{
		Username: p.UserName,
		Password: p.Password,
	}

	// 数据库验证
	if err := database.UserLogin(user); err != nil {
		return 0, "", "", "", err
	}

	// 验证通过后发放token
	aToken, err := jwt.GenToken(user.UserID, user.Username, user.Role, user.Unit)
	return user.UserID, user.Role, user.Unit, aToken, err

}

func GetVerify(num int, ttl time.Duration) (code string) {
	code = genValidateCode(num)
	global.VerifyCode.Ch <- code
	global.VerifyCode.T.Reset(ttl)
	return

}

func Verify(p *model.ParamSignUp) error {
	if err := database.NoUser(); err != nil {
		return err
	} else if p.VerifyCode != global.VerifyCode.Code {
		return database.ErrorVerify
	}
	return nil

}

func genValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
