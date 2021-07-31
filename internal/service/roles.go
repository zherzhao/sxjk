package service

import (
	"encoding/json"
	"webconsole/internal/dao/rbac"
	"webconsole/internal/model"
)

func UpdateRoles(json map[string][]string) error {
	return rbac.RolesUpdate(json)

}

func GetRoles() (*model.RolesResp, error) {
	b, err := rbac.RolesGet()
	if err != nil {
		return nil, err
	}

	r := new(model.RolesResp)
	err = json.Unmarshal(b, r)
	if err != nil {
		return nil, err
	}
	return r, nil

}

func DefaultRoles() error {
	return rbac.RolesDefault()
}
