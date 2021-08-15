package model

// {'datasetNames':["矢量:高速公路"],
// 'getFeatureMode':"SQL",
// 'queryParameter':{'attributeFilter':"路线名称 like '%25绕城%25'"}}
type IServerQueryReq struct {
	DatasetNames   []string `json:"datasetNames"`
	GetFeatureMode string   `json:"getFeatureMode"`
	QueryParameter struct {
		AttributeFilter string `json:"attributeFilter"`
	} `json:"queryParameter"`
}

// {'datasetNames':["矢量:高速公路"],
// 'getFeatureMode':"SQL",
// 'queryParameter':{'name':"高速公路@矢量",
//                   'attributeFilter':"交控单位名称 LIKE '%25绕城%25' OR 路线名称 LIKE '%25绕城%25'",
//                   'orderBy':null},
// 'maxFeatures':-1}
type IServerSearchReq struct {
	DatasetNames   []string `json:"datasetNames"`
	GetFeatureMode string   `json:"getFeatureMode"`
	QueryParameter struct {
		Name            string `json:"name"`
		AttributeFilter string `json:"attributeFilter"`
		OrderBy         string `json:"orderBy"`
	} `json:"queryParameter"`
	MaxFeatures int `jsoin:"maxFeatures"`
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
