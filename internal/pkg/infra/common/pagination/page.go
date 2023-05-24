package pagination

import "math"

type PageResponse[T any] struct {
	Content    []T `json:"items"`
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
}

type Page[T any] interface {
	GetContent() []T
	GetPageable() Pageable
	GetTotalCount() int64
	GetTotalPages() int
	ToResponse() *PageResponse[T]
}

func NewPage[T any](content []T, pageable Pageable, total int64) Page[T] {
	return page[T]{
		content:    content,
		pageable:   pageable,
		totalCount: total,
	}
}

type page[T any] struct {
	content    []T
	pageable   Pageable
	totalCount int64
}

func (p page[T]) GetContent() []T {
	if p.content == nil {
		return []T{}
	}

	return p.content
}

func (p page[T]) GetPageable() Pageable {
	return p.pageable
}

func (p page[T]) GetTotalCount() int64 {
	return p.totalCount
}

func (p page[T]) GetTotalPages() int {
	return int(math.Ceil(float64(p.totalCount) / float64(p.pageable.GetPageSize())))
}

func (p page[T]) ToResponse() *PageResponse[T] {
	return &PageResponse[T]{
		Content:    p.GetContent(),
		TotalCount: int(p.GetTotalCount()),
		TotalPages: p.GetTotalPages(),
	}
}

func MapPage[T any, K any](page Page[T], mapFn func(T) K) Page[K] {
	var content []K
	for _, item := range page.GetContent() {
		content = append(content, mapFn(item))
	}

	return NewPage(content, page.GetPageable(), page.GetTotalCount())
}
