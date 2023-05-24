package pagination

type Sort interface {
	GetSort() string
}

type sort struct {
	sort string
}

func NewSort(s string) Sort {
	if s == "" {
		s = "id"
	}

	return &sort{
		sort: s,
	}
}

func (s *sort) GetSort() string {
	return s.sort
}
