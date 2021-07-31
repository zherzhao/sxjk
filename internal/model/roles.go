package model

type RolesResp struct {
	Root    []string `json:"root"`
	Manager []string `json:"manager"`
	User    []string `json:"user"`
}
