package res

type PageRes struct {
	PageIndex int
	PageSize  int
	Total     int64
	Data      interface{}
}
