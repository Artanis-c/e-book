package req

type PageReq struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

func (p *PageReq) GetOffSet() int {
	return (p.PageIndex - 1) * p.PageSize
}
