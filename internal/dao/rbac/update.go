package rbac

import (
	"webconsole/global"

	"github.com/impact-eintr/WebKits/erbac"
)

func RolesUpdate(json map[string][]string) error {
	tmpRBAC := erbac.NewRBAC()
	tmpPermissions := make(erbac.Permissions)

	for rid, pids := range json {
		role := erbac.NewStdRole(rid)
		for _, pid := range pids {
			_, ok := tmpPermissions[pid]
			if !ok {
				tmpPermissions[pid] = erbac.NewStdPermission(pid)
			}
			role.Assign(tmpPermissions[pid])
		}
		tmpRBAC.Add(role)
	}

	var jsonInher map[string][]string
	// 防止人为修改权限配置文件
	if err := erbac.LoadJson(
		global.RBACSetting.CustomerInherFile, &jsonInher); err != nil {
		return err
	}

	for rid, parents := range jsonInher {
		if err := tmpRBAC.SetParents(rid, parents); err != nil {
			return ErrorInvalidInher
		}
	}
	// 保存权限文件
	err := tmpRBAC.SaveUserRBACWithSort(
		global.RBACSetting.CustomerRoleFile, global.RBACSetting.CustomerInherFile)
	if err != nil {
		return err
	}

	global.Auth.Lock()
	defer global.Auth.Unlock()
	global.Auth.RBAC = tmpRBAC
	global.Auth.Permissions = tmpPermissions

	return nil

}
