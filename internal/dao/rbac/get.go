package rbac

import (
	"io/ioutil"
	"os"
	"webconsole/global"
)

func RolesGet() ([]byte, error) {
	global.Auth.RLock()
	defer global.Auth.RUnlock()

	f, err := os.Open(global.RBACSetting.CustomerRoleFile)
	if err != nil {
		return nil, ErrorRBACNotFound
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil

}
