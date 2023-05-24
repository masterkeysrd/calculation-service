package pagination

type Sort interface {
	ToString() string
}

type sort struct {
	sort string
}

func NewSort(s string) Sort {
	if s == "" {
		s = "id desc"
	}

	return &sort{
		sort: s,
	}
}

func (s *sort) ToString() string {
	return s.sort
}
