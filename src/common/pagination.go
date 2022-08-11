package common

type Pagination struct {
	Number int
	Size   int
}

type Links struct {
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
	Next  *string `json:"next,omitempty"`
	Prev  *string `json:"prev,omitempty"`
	Self  *string `json:"self"`
}

func NewPagination(page int, size int) Pagination {
	return Pagination{page, size}
}
