package req

type SaveBookReq struct {
	BookNo     string  `json:"bookNo"`
	BookName   string  `json:"bookName"`
	BarCode    string  `json:"barCode"`
	CategoryNo string  `json:"category_no"`
	Price      float64 `json:"price"`
	Image      string  `json:"image"`
	Remark     string  `json:"remark"`
}

type BookListReq struct {
	BookName string `json:"bookName"`
	BarCode  string `json:"barCode"`
	UserNo   string
	Page     PageReq
}
