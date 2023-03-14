package res

import "time"

type BookListRes struct {
	BookNo       string    `json:"bookNo"`
	BookName     string    `json:"bookName"`
	CategoryName string    `json:"categoryName"`
	BarCode      string    `json:"barCode"`
	Price        float64   `json:"price"`
	Image        string    `json:"image"`
	Remark       string    `json:"remark"`
	CreateTime   time.Time `json:"createTime"`
}
