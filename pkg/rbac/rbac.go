package rbac

import (
	"log"
	"webconsole/global"

	"github.com/impact-eintr/WebKits/erbac"
)

func Init() error {

	log.Println(global.RBACSetting.RoleFile)
	log.Println(global.RBACSetting.InherFile)

	var err error
	global.RBAC, global.Permissions, err =
		erbac.BuildRBAC(global.RBACSetting.RoleFile, global.RBACSetting.InherFile)
	return err
}
