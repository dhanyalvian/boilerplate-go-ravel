package responses

type Response struct {
	Meta ResponseMeta `json:"meta"`
	Data ResponseData `json:"data,omitempty"`
}

type ResponseList struct {
	Meta ResponseMeta     `json:"meta"`
	Data ResponseListData `json:"data,omitempty"`
}

type ResponseMeta struct {
	CorrId string `json:"correlation_id"`
	Code   string `json:"code"`
	Time   string `json:"time"`
}

type ResponseData struct {
	Record interface{} `json:"record"`
}

type ResponseListData struct {
	Pagination Pagination    `json:"pagination"`
	Records    []interface{} `json:"records"`
}

type EmptyObject struct{}
