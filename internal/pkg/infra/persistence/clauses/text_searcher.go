package clauses

import (
	"fmt"
	"strings"

	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/search"
)

type Fields []string

func (f Fields) Join() string {
	return strings.Join(f, " || ' ' || ")
}

type TextSearcher interface {
	Search(searchable search.Searchable) *SqlBuilder
}

type textSearch struct {
	fields Fields
}

func NewTextSearcher(fields Fields) TextSearcher {
	return &textSearch{
		fields: fields,
	}
}

func (t *textSearch) Search(searchable search.Searchable) *SqlBuilder {
	if searchable.IsEmpty() {
		return NewSqlBuilder("1 = 1")
	}

	sql := t.buildSQL(searchable)
	return NewSqlBuilder(sql)
}

func (t *textSearch) buildSQL(searchable search.Searchable) string {
	fields := t.fields.Join()
	terms := searchable.Terms()
	search := strings.Join(terms, ":* | ") + ":*"

	return fmt.Sprintf("to_tsvector(%s) @@ to_tsquery('%s')", fields, search)
}
