package clauses

import "gorm.io/gorm/clause"

type SqlBuilder struct {
	sql string
}

func NewSqlBuilder(sql string) *SqlBuilder {
	return &SqlBuilder{
		sql: sql,
	}
}

func (t *SqlBuilder) Build(build clause.Builder) {
	build.WriteString(t.sql)
}
