package rbac

import (
	"webconsole/global"

	"github.com/impact-eintr/WebKits/erbac"
)

func RolesDefault() (err error) {
	global.Auth.Lock()
	defer global.Auth.Unlock()

	global.Auth.RBAC, global.Auth.Permissions, err = erbac.BuildRBAC(
		global.RBACSetting.DefaultRoleFile, global.RBACSetting.DefaultInherFile)
	if err != nil {
		return
	}

	err = global.Auth.RBAC.SaveUserRBACWithSort(
		global.RBACSetting.CustomerRoleFile, global.RBACSetting.CustomerInherFile)
	if err != nil {
		return
	}
	return

}
