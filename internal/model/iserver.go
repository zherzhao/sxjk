package model

type IServerReq struct {
	DatasetNames   []string `json:"datasetNames"`
	GetFeatureMode string   `json:"getFeatureMode"`
	QueryParameter struct {
		AttributeFilter string `json:"attributeFilter"`
	} `json:"queryParameter"`
}

type IServerResp struct {
	PostResultType      string `json:"postResultType"`
	NewResourceID       string `json:"newResourceID"`
	Succeed             bool   `json:"succeed"`
	NewResourceLocation string `json:"newResourceLocation"`
}

type IServerFeatures struct {
	Features       []Test   `json:"features"`
	FeatureUriList []string `json:"featureUriList"`
	TotalCount     int      `json:"totalCount"`
	FeatureCount   int      `json:"featureCount"`
}

type Test struct {
	StringID    []string `json:"stringID"`
	FieldNames  []string `json:"fieldNames"`
	Geometry    Geo      `json:"geometry"`
	FieldValues []string `json:"fieldValues"`
	ID          int      `json:"ID"`
}

type Geo struct {
	Center struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"center"`
	Parts       []int    `json:"parts"`
	Style       []string `json:"style"`
	PrjCoordSys []string `json:"prjCoordSys"`
	Id          int      `json:"id"`
	Type        string   `json:"type"`
	PartTopo    []string `json:"partTopo"`
	Points      []struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"points"`
}
