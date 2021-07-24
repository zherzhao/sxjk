package rbac

import (
	"errors"
	"os"
	"webconsole/global"

	"github.com/impact-eintr/WebKits/erbac"
)

func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Init() error {

	var err error

	if exists(global.RBACSetting.CustomerInherFile) &&
		exists(global.RBACSetting.CustomerRoleFile) {

		global.Auth.RBAC, global.Auth.Permissions, err =
			erbac.BuildRBAC(global.RBACSetting.CustomerRoleFile,
				global.RBACSetting.CustomerInherFile)

	} else if exists(global.RBACSetting.DefaultInherFile) &&
		exists(global.RBACSetting.DefaultRoleFile) {

		global.Auth.RBAC, global.Auth.Permissions, err =
			erbac.BuildRBAC(global.RBACSetting.DefaultRoleFile,
				global.RBACSetting.DefaultInherFile)

	} else {
		err = errors.New("权限配置文件不存在")
	}
	return err

}
