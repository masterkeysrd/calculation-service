package common

type ListRequest struct {
	Search string `form:"search"`
	Page   int    `form:"page"`
	Size   int    `form:"size"`
	Sort   string `form:"sort"`
}

func (r *ListRequest) GetSearch() string {
	return r.Search
}

func (r *ListRequest) GetPagination() PaginationRequest {
	return PaginationRequest{
		Number: r.Page,
		Size:   r.Size,
	}
}

func (r *ListRequest) GetSort() string {
	return r.Sort
}
