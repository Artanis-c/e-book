package res

type PageRes struct {
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	Total     int64       `json:"total"`
	Records   interface{} `json:"records"`
}
