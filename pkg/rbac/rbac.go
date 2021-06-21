package rbac

import (
	"log"
	"webconsole/global"

	"github.com/impact-eintr/WebKits/erbac"
)

func Init() error {

	log.Println(global.RBACSetting.RoleFile)
	log.Println(global.RBACSetting.InherFile)

	global.RBAC, global.Permissions =
		erbac.BuildRBAC(global.RBACSetting.RoleFile, global.RBACSetting.InherFile)
	return nil
}
