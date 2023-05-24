package pagination

type PageableRequest struct {
	Page int    `json:"page" form:"page"`
	Size int    `json:"size" form:"size"`
	Sort string `json:"sort" form:"sort"`
}

type Pageable interface {
	GetSort() Sort
	GetOffset() int
	GetPageNumber() int
	GetPageSize() int
}

type pageable struct {
	pageNumber int
	pageSize   int
	sort       Sort
}

func NewPageable(page PageableRequest) Pageable {
	pageNumber := page.Page
	pageSize := page.Size

	if pageNumber < 1 {
		pageNumber = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	return &pageable{
		pageNumber: pageNumber,
		pageSize:   pageSize,
		sort:       NewSort(page.Sort),
	}
}

func (p *pageable) GetSort() Sort {
	return p.sort
}

func (p *pageable) GetOffset() int {
	return (p.GetPageNumber() - 1) * p.GetPageSize()
}

func (p *pageable) GetPageNumber() int {
	return p.pageNumber
}

func (p *pageable) GetPageSize() int {
	return p.pageSize
}
