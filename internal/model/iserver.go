package model

type IServerReq struct {
	DatasetNames   []string `json:"datasetNames"`
	GetFeatureMode string   `json:"getFeatureMode"`
	QueryParameter struct {
		AttributeFilter string `json:"attributeFilter"`
	} `json:"queryParameter"`
}

type IServerResp struct{}
